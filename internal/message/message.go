package message

// Failure 错误时返回结构
type Failure struct {
	Code       int    `json:"code"`        // 业务码
	Message    string `json:"message"`     // 描述信息
	ErrorStack string `json:"error-stack"` // 描述信息
}

const (
	/* 服务级错误码 */
	ServerError        = 10101
	TooManyRequests    = 10102
	ParamBindError     = 10103
	AuthorizationError = 10104
	CallHTTPError      = 10105
)

var codeText = map[int]string{
	ServerError:        "Internal Server Error",
	TooManyRequests:    "Too Many Requests",
	ParamBindError:     "参数信息有误",
	AuthorizationError: "签名信息有误",
	CallHTTPError:      "调用第三方 HTTP 接口失败",
}

func Text(code int) string {
	return codeText[code]
}
