package erron

type Code int

// 错误码列表
const (
	Success         Code = 200
	Failed          Code = 10000
	UnLogin         Code = 401
	ErrInternal     Code = 10003
	ErrRequestParam Code = 10004
	ErrAccountInfo  Code = 10005
)

// 错误码说明
var errCodeMap = map[Code]string{
	Success:         "请求成功",
	Failed:          "请求失败",
	UnLogin:         "未登录",
	ErrInternal:     "内部错误或异常",
	ErrRequestParam: "请求参数错误",
	ErrAccountInfo:  "账号信息有误",
}

func (c Code) Msg() string {
	if v, ok := errCodeMap[c]; ok {
		return v
	}
	return ``
}
