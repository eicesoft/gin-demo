package controller

import (
	"eicesoft/web-demo/internal/service/config_service"
	"eicesoft/web-demo/pkg/errno"
	"eicesoft/web-demo/pkg/message"
	"fmt"
	"net/http"

	"eicesoft/web-demo/pkg/core"
	"eicesoft/web-demo/pkg/db"
	"eicesoft/web-demo/pkg/mux"

	"go.uber.org/zap"
)

const GroupRouterName = "/config"

//var _ ConfigHandler = (*configHandler)(nil)

// ConfigHandler 配置控制器
type ConfigHandler interface {
	RegistryRouter(r *mux.Resource)
	List() *core.RouteInfo
	Store() *core.RouteInfo
}

// configHandler struct
type configHandler struct {
	logger        *zap.Logger
	db            db.Repo
	configService config_service.ConfigService
}

// NewConfigHandler 返回Handler
func NewConfigHandler(logger *zap.Logger, db db.Repo) *configHandler {
	return &configHandler{
		logger:        logger,
		db:            db,
		configService: config_service.NewConfigService(db),
	}
}

// RegistryRouter 注册路由
func (h *configHandler) RegistryRouter(r *mux.Resource) {
	config := r.Mux.Group(GroupRouterName, core.WrapAuthHandler(r.Middles.Jwt))
	config.WrapRouters(
		h.List(),
		h.Store(),
	)
}

// List 列出配置
func (h *configHandler) List() *core.RouteInfo {
	return &core.RouteInfo{
		Method: core.GET,
		Path:   "list",
		Closure: func(c core.Context) {
			list := h.configService.List(c)
			c.Payload(list)
		},
	}
}

type storeRequest struct {
	Id    int    `form:"id"`
	Key   string `form:"key"`
	Value string `form:"value"`
	Title string `form:"title"`
	Type  int    `form:"type"`
}

func (h *configHandler) Store() *core.RouteInfo {
	return &core.RouteInfo{
		Method: core.POST,
		Path:   "store",
		Closure: func(c core.Context) {
			req := new(storeRequest)
			if err := c.ShouldBindForm(req); err != nil {
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					message.ParamBindError,
					message.Get().Text(message.ParamBindError),
					err).WithErr(err),
				)
				return
			}

			fmt.Println(req.Title)
			c.Payload(req)
		},
	}
}
