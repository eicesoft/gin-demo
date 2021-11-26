package middleware

import (
	"eicesoft/web-demo/config"
	"eicesoft/web-demo/pkg/core"
	"eicesoft/web-demo/pkg/errno"
	"eicesoft/web-demo/pkg/message"
	"eicesoft/web-demo/pkg/token"
	"net/http"

	"github.com/pkg/errors"
)

func (m *middleware) Jwt(ctx core.Context) (userId int64, err errno.Error) {
	auth := ctx.GetHeader("Authorization")
	if auth == "" {
		e := errors.New("Header 中缺少 Authorization 参数")
		err = errno.NewError(
			http.StatusUnauthorized,
			message.AuthorizationError,
			message.Get().Text(message.AuthorizationError),
			e).WithErr(e)

		return
	}

	cfg := config.Get().JWT
	claims, errParse := token.New(cfg.Secret).JwtParse(auth)
	if errParse != nil {
		err = errno.NewError(
			http.StatusUnauthorized,
			message.AuthorizationError,
			message.Get().Text(message.AuthorizationError),
			errParse).WithErr(errParse)

		return
	}

	userId = claims.UserID
	if userId <= 0 {
		e := errors.New("claims.UserID <= 0 ")
		err = errno.NewError(
			http.StatusUnauthorized,
			message.AuthorizationError,
			message.Get().Text(message.AuthorizationError),
			e).WithErr(e)

		return
	}
	return
}
