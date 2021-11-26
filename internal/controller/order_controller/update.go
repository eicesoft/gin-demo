package order_controller

import (
	"eicesoft/web-demo/internal/service/order_service"
	"eicesoft/web-demo/pkg/core"
	"eicesoft/web-demo/pkg/errno"
	"eicesoft/web-demo/pkg/message"
	"net/http"
)

type updateResponse struct {
	Id int32 `json:"id"` // 主键ID
}

func (h *handler) Update() *core.RouteInfo {
	return &core.RouteInfo{
		Method: core.POST,
		Path:   "update",
		Closure: func(c core.Context) {
			req := new(order_service.OrderInfo)
			if err := c.ShouldBindForm(req); err != nil {
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					message.ParamBindError,
					message.Get().Text(message.ParamBindError),
					err).WithErr(err),
				)
				return
			}

			err := h.orderService.Update(c, req)
			if err != nil {
				panic(err)
			}

			res := new(updateResponse)
			res.Id = 1
			c.Payload(res)
		},
	}
}
