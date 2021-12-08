package controller

import (
	"eicesoft/web-demo/pkg/core"
	"eicesoft/web-demo/pkg/mux"
)

type RouteInterface interface {
	Params() (string, core.HandlerFunc)
}

// IdRequest Id基本请求
type IdRequest struct {
	Id int `form:"id"`
}

// PageRequest 分页基本请求
type PageRequest struct {
	Page     int `form:"page"`
	PageSize int `form:"page_size"`
}

func RegistryRouters(r *mux.Resource) {
	NewConfigHandler(r.Mux.GetLogger(), r.Mux.GetDB()).RegistryRouter(r)
}
