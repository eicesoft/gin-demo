package auth_controller

import (
	"eicesoft/web-demo/config"
	"eicesoft/web-demo/pkg/core"
	"eicesoft/web-demo/pkg/errno"
	"eicesoft/web-demo/pkg/message"
	"eicesoft/web-demo/pkg/token"
	"math/rand"
	"net/http"
	"time"
)

type authRequest struct {
	UserName string `form:"username" binding:"required"` // 用户名
	Password string `form:"password" binding:"required"` // 密码
}

type authResponse struct {
	Authorization string `json:"authorization"` // 签名
	ExpireTime    int64  `json:"expire_time"`   // 过期时间
}

func (h *handler) Get() *core.RouteInfo {
	return &core.RouteInfo{
		Method: core.GET,
		Path:   "get",
		Closure: func(c core.Context) {
			req := new(authRequest)
			//res := new(detailResponse)
			if err := c.ShouldBindQuery(req); err != nil {
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					message.ParamBindError,
					message.Get().Text(message.ParamBindError),
					err).WithErr(err),
				)
				return
			}

			cfg := config.Get().JWT
			tokenString, err := token.New(cfg.Secret).JwtSign(rand.Int63n(10000), time.Hour*cfg.ExpireDuration)
			if err != nil {
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					message.AuthorizationError,
					message.Get().Text(message.AuthorizationError),
					err).WithErr(err),
				)
				return
			}

			res := new(authResponse)
			res.Authorization = tokenString
			res.ExpireTime = time.Now().Add(time.Hour * cfg.ExpireDuration).Unix()

			c.Payload(res)
		},
	}
}
