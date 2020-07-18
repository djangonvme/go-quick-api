package erron

import "fmt"

// 错误接口
type E interface {
	Error() string
	Code() Code
	Msg() string
}

// 错误信息
type errInfo struct {
	code Code
	msg  string
}

func (e *errInfo) Error() string {
	return e.msg
}

func (e *errInfo) Code() Code {
	return e.code
}

func (e *errInfo) Msg() string {
	return e.msg
}

func textToMsg(text ...string) string {
	var msg string
	if len(text) > 0 {
		msg = fmt.Sprintf(text[0], text[1:])
	}
	return msg
}

// 新建一个错误 code 必填，text可选
func New(code Code, text ...string) E {
	return &errInfo{code: code, msg: textToMsg(text...)}
}

// 通用型失败 code=10000
func Fail(text ...string) E {
	return &errInfo{code: Failed, msg: textToMsg(text...)}
}

// 从另一个错误接口返回
func FailBy(err error) E {
	return &errInfo{code: Failed, msg: err.Error()}
}

// 内部错误，无关业务逻辑
func Inner(text ...string) E {
	return &errInfo{code: ErrInternal, msg: textToMsg(text...)}
}

// 仅在合适情况使用，直接将错误抛出为异常。
func Try(err error) {
	if err != nil {
		panic(err)
	}
}
