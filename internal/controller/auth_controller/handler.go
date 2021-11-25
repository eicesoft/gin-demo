package auth_controller

import (
	"eicesoft/web-demo/pkg/core"
	"eicesoft/web-demo/pkg/db"
	"eicesoft/web-demo/pkg/mux"
	"go.uber.org/zap"
)

const GroupRouterName = "/auth"

var _ Handler = (*handler)(nil)

// Handler 用户控制器接口
type Handler interface {
	RegistryRouter(r *mux.Resource)
	Get() *core.RouteInfo
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

func (h *handler) RegistryRouter(r *mux.Resource) {
	auth := r.Mux.Group(GroupRouterName)
	auth.WrapRouters(h.Get())
}
