package user_controller

import (
	"eicesoft/web-demo/internal/message"
	"eicesoft/web-demo/pkg/core"
	"eicesoft/web-demo/pkg/errno"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type detailRequest struct {
	UserName string `uri:"username"` // 用户名
}

type detailResponse struct {
	Id       int64       `json:"id"`        // 用户主键ID
	UserName string      `json:"user_name"` // 用户名
	NickName string      `json:"nick_name"` // 昵称
	Data     interface{} `json:"data"`
}

func (h *handler) Test() (string, core.HandlerFunc) {
	return "test", func(c core.Context) {
		c.Payload(gin.H{
			"message": "这个是一个Gin.H消息",
			"data":    c.Header(),
		})
	}
}

// Detail 用户详情
// @Summary 用户详情
// @Description 用户详情
// @Tags User
// @Accept  json
// @Produce  json
// @Param username path string true "用户名"
// @Success 200 {object} detailResponse
// @Failure 400 {object} message.Failure
// @Failure 401 {object} message.Failure
// @Router /user/get/{username} [get]
func (h *handler) Detail() (string, core.HandlerFunc) {
	return "get/:username", func(c core.Context) {
		req := new(detailRequest)
		res := new(detailResponse)
		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				message.ParamBindError,
				message.Text(message.ParamBindError),
				err).WithErr(err),
			)
			return
		}

		u := h.userService.Get()

		//u := user.User{}
		//h.db.GetDbR().WithContext(c.RequestContext()).First(&u)

		if req.UserName != "sdg" {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				message.ParamBindError,
				message.Text(message.ParamBindError),
				errors.New("未知错误")),
			)
			return
		}

		res.Id = c.UserID()
		res.UserName = req.UserName
		res.NickName = req.UserName + "_nick"
		res.Data = u

		c.Payload(res)
	}
}
