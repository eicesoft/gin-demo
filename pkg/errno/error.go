package errno

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var _ Error = (*err)(nil)

type Error interface {
	p()
	// WithErr 设置错误信息
	WithErr(err error) Error
	// GetBusinessCode 获取 Business Code
	GetBusinessCode() int
	// GetHttpCode 获取 HTTP Code
	GetHttpCode() int
	// GetMsg 获取 Msg
	GetMsg() string
	// GetBusinessMsg 获得异常消息
	GetBusinessMsg() string
	// ToString 返回 JSON 格式的错误详情
	ToString() string
}

type err struct {
	HttpCode        int
	BusinessCode    int
	Message         string
	Err             error
	BusinessMessage string
}

func NewError(httpCode, businessCode int, msg string, businessErr error) Error {
	e := &err{
		HttpCode:        httpCode,
		BusinessCode:    businessCode,
		Message:         msg,
		BusinessMessage: businessErr.Error(),
	}

	if businessErr != nil {
		e.WithErr(businessErr)
	}

	return e
}

func (e *err) p() {}

func (e *err) GetBusinessMsg() string {
	return e.BusinessMessage
}

func (e *err) WithErr(err error) Error {
	e.Err = errors.WithStack(err)
	return e
}

func (e *err) GetHttpCode() int {
	return e.HttpCode
}

func (e *err) GetBusinessCode() int {
	return e.BusinessCode
}

func (e *err) GetMsg() string {
	return e.Message
}

// ToString 返回 JSON 格式的错误详情
func (e *err) ToString() string {
	err := &struct {
		HttpCode        int    `json:"code"`
		BusinessCode    int    `json:"error_code"`
		Message         string `json:"message"`
		BusinessMessage string `json:"business-message"`
	}{
		HttpCode:        e.HttpCode,
		BusinessCode:    e.BusinessCode,
		Message:         e.Message,
		BusinessMessage: e.BusinessMessage,
	}

	raw, _ := json.Marshal(err)
	return string(raw)
}
