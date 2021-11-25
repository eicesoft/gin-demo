package core

import (
	"eicesoft/web-demo/pkg/color"
	"eicesoft/web-demo/pkg/env"
	"fmt"
	"github.com/gin-gonic/gin"
)

var _ IRoute = (*router)(nil)

type RequestMethod int

const (
	ANY = iota
	GET
	POST
	DELETE
	PATCH
	PUT
	OPTIONS
	HEAD
)

// IRoute 基本包装gin的IRoutes
type IRoute interface {
	Any(string, ...HandlerFunc)
	GET(string, ...HandlerFunc)
	POST(string, ...HandlerFunc)
	DELETE(string, ...HandlerFunc)
	PATCH(string, ...HandlerFunc)
	PUT(string, ...HandlerFunc)
	OPTIONS(string, ...HandlerFunc)
	HEAD(string, ...HandlerFunc)
}

type router struct {
	group *gin.RouterGroup
}

type RouteInfo struct {
	Method  RequestMethod
	Path    string
	Closure HandlerFunc
}

func (r *RouteInfo) Params() (string, HandlerFunc) {
	return r.Path, r.Closure
}

// WrapRouters 包装路由信息
func (r *router) WrapRouters(routes ...*RouteInfo) {
	for i := range routes {
		route := *routes[i]

		if !env.Get().IsProd() {
			fmt.Println(color.Green(fmt.Sprintf("* register route: %s/%s", r.group.BasePath(), route.Path)))
		}

		switch route.Method {
		case GET:
			r.GET(route.Params())
		case POST:
			r.POST(route.Params())
		case DELETE:
			r.DELETE(route.Params())
		case HEAD:
			r.HEAD(route.Params())
		case PATCH:
			r.PATCH(route.Params())
		case OPTIONS:
			r.OPTIONS(route.Params())
		case PUT:
			r.PUT(route.Params())
		case ANY:
			r.Any(route.Params())
		default:
			r.Any(route.Params())
		}
	}
}

func (r *router) Group(relativePath string, handlers ...HandlerFunc) RouterGroup {
	group := r.group.Group(relativePath, wrapHandlers(handlers...)...)
	return &router{group: group}
}

func (r *router) Any(relativePath string, handlers ...HandlerFunc) {
	r.group.Any(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) GET(relativePath string, handlers ...HandlerFunc) {
	r.group.GET(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) POST(relativePath string, handlers ...HandlerFunc) {
	r.group.POST(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) DELETE(relativePath string, handlers ...HandlerFunc) {
	r.group.DELETE(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) PATCH(relativePath string, handlers ...HandlerFunc) {
	r.group.PATCH(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) PUT(relativePath string, handlers ...HandlerFunc) {
	r.group.PUT(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) OPTIONS(relativePath string, handlers ...HandlerFunc) {
	r.group.OPTIONS(relativePath, wrapHandlers(handlers...)...)
}

func (r *router) HEAD(relativePath string, handlers ...HandlerFunc) {
	r.group.HEAD(relativePath, wrapHandlers(handlers...)...)
}
