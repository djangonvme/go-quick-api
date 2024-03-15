package erron

type ErrCode int

// 错误码列表
const (
	Success         ErrCode = 200
	Failed          ErrCode = 1000
	UnLogin         ErrCode = 401
	ErrInternal     ErrCode = 1003
	ErrRequestParam ErrCode = 1004
	ErrAccountInfo  ErrCode = 1005
	PanicCode       ErrCode = 1006
)

// 错误码说明
var errCodeMap = map[ErrCode]string{
	Success:         "请求成功",
	Failed:          "请求失败",
	UnLogin:         "未登录",
	ErrInternal:     "内部错误或异常",
	ErrRequestParam: "请求参数错误",
	ErrAccountInfo:  "账号信息有误",
}

func (c ErrCode) Text() string {
	if v, ok := errCodeMap[c]; ok {
		return v
	}
	return ``
}
