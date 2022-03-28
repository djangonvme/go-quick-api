package erron

import (
	"strings"
)

type LogIF interface {
	write(err error)
}

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
	return strings.Join(text, ";")
}

func New(code Code, text ...string) E {
	return &errInfo{code: code, msg: textToMsg(text...)}
}

func Inner(text ...string) E {
	return &errInfo{code: ErrInternal, msg: textToMsg(text...)}
}
