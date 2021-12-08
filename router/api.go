package router

import (
	"eicesoft/web-demo/internal/controller"
	"eicesoft/web-demo/internal/controller/auth_controller"
	"eicesoft/web-demo/internal/controller/order_controller"
	"eicesoft/web-demo/internal/controller/user_controller"
	"eicesoft/web-demo/pkg/mux"
)

// 设置Api路由
func setApiRouter(r *mux.Resource) {
	user_controller.New(r.Mux.GetLogger(), r.Mux.GetDB()).RegistryRouter(r)
	auth_controller.New(r.Mux.GetLogger(), r.Mux.GetDB()).RegistryRouter(r)
	order_controller.New(r.Mux.GetLogger(), r.Mux.GetDB()).RegistryRouter(r)
	controller.RegistryRouters(r)
}
