package order_controller

import (
	"eicesoft/web-demo/pkg/core"
)

//type orderListResponse struct {
//	Data []*orders.Order `json:"data"`
//}

// List 订单列表
// @Summary 订单列表
// @Description 订单列表
// @Tags Order
// @Accept  json
// @Produce  json
// @Param Authorization header string true "验证Token"
// @Success 200 {object} orderListResponse
// @Failure 400 {object} message.Failure
// @Failure 401 {object} message.Failure
// @Router /order/list [get]
func (h *handler) List() *core.RouteInfo {
	return &core.RouteInfo{
		Method: core.GET,
		Path: "list",
		Closure: func(c core.Context) {
			list := h.orderService.List(c)
			//res := new(orderListResponse)
			//res.Code = 200
			//res.Data = list
			c.Payload(list)
		},
	}
}
