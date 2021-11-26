
package config_controller

import (
	"eicesoft/web-demo/pkg/core"
	"eicesoft/web-demo/pkg/db"
	"eicesoft/web-demo/pkg/mux"
	"go.uber.org/zap"
)

const GroupRouterName = "/config"

var _ Handler = (*handler)(nil)

// Handler
type Handler interface {
	RegistryRouter(r *mux.Resource)
	Test() *core.RouteInfo	//Demo
}

type handler struct {
	logger *zap.Logger
	db     db.Repo
}

func New(logger *zap.Logger, db db.Repo) Handler {
	return &handler{
		logger: logger,
		db:     db,
	}
}

func (h *handler) Test() *core.RouteInfo {
	return &core.RouteInfo{
		Method: core.GET,
		Path:   "test",
		Closure: func(c core.Context) {
			c.Payload("")
		},
	}
}

func (h *handler) RegistryRouter(r *mux.Resource) {
	auth := r.Mux.Group(GroupRouterName)
	auth.WrapRouters(h.Test())
}