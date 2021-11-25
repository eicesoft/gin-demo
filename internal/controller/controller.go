package controller

import (
	"eicesoft/web-demo/pkg/core"
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
