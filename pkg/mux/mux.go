package mux

import (
	"eicesoft/web-demo/internal/middleware"
	"eicesoft/web-demo/pkg/core"
)

type Resource struct {
	Mux     core.Mux
	Middles middleware.Middleware
}

type ResourceInterface interface {
}
