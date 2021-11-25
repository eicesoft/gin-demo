package core

import (
	"eicesoft/web-demo/config"
	_ "eicesoft/web-demo/docs"
	"eicesoft/web-demo/internal/message"
	"eicesoft/web-demo/pkg/color"
	"eicesoft/web-demo/pkg/env"
	"eicesoft/web-demo/pkg/errno"
	"eicesoft/web-demo/pkg/trace"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"net/http"
	"net/url"
	"reflect"
	"runtime/debug"
	"time"
)

const (
	ServerHeader = "Server"
	ServerName   = "Gee Server"
	MaxBurstSize = 10000
)

// BaseResponse 基本响应
type BaseResponse struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// withoutTracePaths 不追踪处理的path
var withoutTracePaths = map[string]bool{
	"/metrics":                  true,
	"/debug/pprof/":             true,
	"/debug/pprof/cmdline":      true,
	"/debug/pprof/profile":      true,
	"/debug/pprof/symbol":       true,
	"/debug/pprof/trace":        true,
	"/debug/pprof/allocs":       true,
	"/debug/pprof/block":        true,
	"/debug/pprof/goroutine":    true,
	"/debug/pprof/heap":         true,
	"/debug/pprof/mutex":        true,
	"/debug/pprof/threadcreate": true,
	"/favicon.ico":              true,
	"/system/health":            true,
}

// StructCopy 从value复制结构数据到 binding 中
func StructCopy(binding interface{}, value interface{}) {
	bVal := reflect.ValueOf(binding).Elem()
	vVal := reflect.ValueOf(value).Elem()
	vTypeOfT := vVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		name := vTypeOfT.Field(i).Name
		if ok := bVal.FieldByName(name).IsValid(); ok { //目标结构中存在字段
			bVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface())) //
		}
	}
}

func wrapHandlers(handlers ...HandlerFunc) []gin.HandlerFunc {
	funcs := make([]gin.HandlerFunc, len(handlers))
	for i, handler := range handlers {
		handler := handler
		funcs[i] = func(c *gin.Context) {
			ctx := newContext(c)
			defer releaseContext(ctx)

			handler(ctx)
		}
	}

	return funcs
}

// WrapAuthHandler 用来处理 Auth 的入口，在之后的handler中只需 ctx.UserID()
func WrapAuthHandler(handler func(Context) (userID int64, err errno.Error)) HandlerFunc {
	return func(ctx Context) {
		userID, err := handler(ctx)
		if err != nil {
			ctx.AbortWithError(err)
			return
		}
		ctx.setUserID(userID)
	}
}

type RouterGroup interface {
	IRoute
	Group(string, ...HandlerFunc) RouterGroup
	WrapRouters(routes ...*RouteInfo)
}

func Alias(path string) HandlerFunc {
	return func(ctx Context) {
		ctx.setAlias(path)
	}
}

var _ Mux = (*mux)(nil)

type Mux interface {
	http.Handler
	Group(relativePath string, handlers ...HandlerFunc) RouterGroup
}

type mux struct {
	engine *gin.Engine
}

func (m *mux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	m.engine.ServeHTTP(w, req)
}

func (m *mux) Group(relativePath string, handlers ...HandlerFunc) RouterGroup {
	return &router{
		group: m.engine.Group(relativePath, wrapHandlers(handlers...)...),
	}
}

func DisableTrace(ctx Context) {
	ctx.disableTrace()
}

func New(logger *zap.Logger) (Mux, error) {
	ui := `
 ██████╗ ██╗    ██╗███████╗██████╗      █████╗ ██████╗ ██╗
██╔════╝ ██║    ██║██╔════╝██╔══██╗    ██╔══██╗██╔══██╗██║
██║  ███╗██║ █╗ ██║█████╗  ██████╔╝    ███████║██████╔╝██║
██║   ██║██║███╗██║██╔══╝  ██╔══██╗    ██╔══██║██╔═══╝ ██║
╚██████╔╝╚███╔███╔╝███████╗██████╔╝    ██║  ██║██║     ██║
 ╚═════╝  ╚══╝╚══╝ ╚══════╝╚═════╝     ╚═╝  ╚═╝╚═╝     ╚═╝`
	fmt.Println(color.Green(ui))
	gin.SetMode(gin.ReleaseMode)
	gin.DisableBindValidation()

	mux := &mux{
		engine: gin.New(),
	}

	fmt.Println(color.Green(fmt.Sprintf("* listen port: %s", config.Get().Server.Port)))
	fmt.Println(color.Green(fmt.Sprintf("* run env: %s", env.Get().Value())))

	if !env.Get().IsProd() {
		pprof.Register(mux.engine) // register pprof to gin
		fmt.Println(color.Green("* register pprof"))
	}

	if !env.Get().IsProd() {
		mux.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // register swagger
		fmt.Println(color.Green("* register swagger router"))
	}

	if env.Get().IsProd() { //开启限流
		limiter := rate.NewLimiter(rate.Every(time.Second*1), MaxBurstSize)
		mux.engine.Use(func(ctx *gin.Context) {
			context := newContext(ctx)
			defer releaseContext(context)

			if !limiter.Allow() {
				context.AbortWithError(errno.NewError(
					http.StatusTooManyRequests,
					message.TooManyRequests,
					message.Text(message.TooManyRequests),
					fmt.Errorf("")),
				)
				return
			}

			ctx.Next()
		})
	}

	// recover两次，防止处理时发生panic，尤其是在OnPanicNotify中。
	mux.engine.Use(func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//panic(err)
				logger.Error("got panic", zap.String("panic", fmt.Sprintf("%+v", err)), zap.String("stack", string(debug.Stack())))
			}
		}()

		ctx.Next()
	})

	if config.Get().Server.Cors {
		fmt.Println(color.Green("* register cors middleware"))
		mux.engine.Use(cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{
				http.MethodHead,
				http.MethodGet,
				http.MethodPost,
				http.MethodPut,
				http.MethodPatch,
				http.MethodDelete,
			},
			AllowedHeaders:     []string{"*"},
			AllowCredentials:   true,
			OptionsPassthrough: true,
		}))
	}

	mux.engine.Use(func(ctx *gin.Context) {
		coreProcess(ctx, logger)
	})

	mux.engine.NoMethod(wrapHandlers(DisableTrace)...)
	mux.engine.NoRoute(wrapHandlers(DisableTrace)...)
	system := mux.Group("/system")
	{
		system.GET("/health", func(ctx Context) {
			resp := &struct {
				Timestamp   time.Time `json:"timestamp"`
				Environment string    `json:"environment"`
				Host        string    `json:"host"`
				Status      string    `json:"status"`
			}{
				Timestamp:   time.Now(),
				Environment: env.Get().Value(),
				Host:        ctx.Host(),
				Status:      "ok",
			}
			ctx.Payload(resp)
		})
	}

	return mux, nil
}

func coreProcess(ctx *gin.Context, logger *zap.Logger) {
	ts := time.Now()
	//核心处理
	context := newContext(ctx)
	defer releaseContext(context)

	context.init()
	context.setLogger(logger)

	if !withoutTracePaths[ctx.Request.URL.Path] {
		//trace id 前端Header传递该值, 方便调试
		if traceId := context.GetHeader(trace.Header); traceId != "" {
			context.setTrace(trace.New(traceId))
		} else {
			context.setTrace(trace.New(""))
		}
	}

	defer func() {
		if err := recover(); err != nil {
			logger.Error("Http request Error:", zap.Any("err", err))

			context.AbortWithError(errno.NewError(
				http.StatusInternalServerError,
				message.ServerError,
				message.Text(message.ServerError),
				fmt.Errorf("%+v", err)),
			)
		}

		if ctx.Writer.Status() == http.StatusNotFound {
			return
		}
		var (
			response        interface{}
			businessCode    int
			businessCodeMsg string
			errorStack      string
			abortErr        error
			//traceId         string
		)

		context.SetHeader(ServerHeader, ServerName)

		if ctx.IsAborted() { //
			if err := context.abortError(); err != nil {
				response = err
				businessCode = err.GetBusinessCode()
				businessCodeMsg = err.GetMsg()
				errorStack = err.GetMsg()

				if env.Get().IsProd() {
					errorStack = ""
				} else {
					errorStack = err.GetBusinessMsg()
				}

				if x := context.Trace(); x != nil {
					context.SetHeader(trace.Header, x.ID())
					//traceId = x.ID()
				}

				ctx.JSON(err.GetHttpCode(), &message.Failure{
					Code:       businessCode,
					Message:    businessCodeMsg,
					ErrorStack: errorStack,
				})
			}
		} else {
			response = context.getPayload()
			if response != nil {
				if x := context.Trace(); x != nil {
					context.SetHeader(trace.Header, x.ID()) //设置Trace Id
					//traceId = x.ID()
				}
				res := new(BaseResponse)
				res.Code = 200
				res.Data = response
				ctx.JSON(http.StatusOK, res)
			}
		}

		var t *trace.Trace
		if x := context.Trace(); x != nil {
			t = x.(*trace.Trace)
		} else {
			return
		}
		decodedURL, _ := url.QueryUnescape(ctx.Request.URL.RequestURI())
		t.WithRequest(&trace.Request{
			TTL:        "un-limit",
			Method:     ctx.Request.Method,
			DecodedURL: decodedURL,
			//Header:     ctx.Request.Header,
			Body: string(context.RawData()),
		})

		var responseBody interface{}

		if response != nil {
			responseBody = response
		}

		t.WithResponse(&trace.Response{
			Header:          ctx.Writer.Header(),
			HttpCode:        ctx.Writer.Status(),
			HttpCodeMsg:     http.StatusText(ctx.Writer.Status()),
			BusinessCode:    businessCode,
			BusinessCodeMsg: businessCodeMsg,
			Body:            responseBody,
			CostSeconds:     time.Since(ts).Seconds(),
		})

		t.Success = !ctx.IsAborted() && ctx.Writer.Status() == http.StatusOK
		t.CostSeconds = time.Since(ts).Seconds()

		logger.Debug("router-interceptor",
			zap.Any("method", ctx.Request.Method),
			zap.Any("path", decodedURL),
			zap.Any("http_code", ctx.Writer.Status()),
			zap.Any("business_code", businessCode),
			zap.Any("success", t.Success),
			zap.Any("cost_seconds", t.CostSeconds),
			zap.Any("trace_id", t.Identifier),
			zap.Any("trace_info", t),
			zap.Error(abortErr),
		)
	}()
	ctx.Next()
}
