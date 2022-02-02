package erron

import (
	"fmt"
	"strings"
)

type LogIF interface {
	write(err error)
}

var logger LogIF

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

func Fail(text ...string) E {
	return &errInfo{code: Failed, msg: textToMsg(text...)}
}

func FailBy(err error) E {
	return &errInfo{code: Failed, msg: err.Error()}
}

func Inner(text ...string) E {
	return &errInfo{code: ErrInternal, msg: textToMsg(text...)}
}

func Try(err error) {
	if err != nil {
		panic(err)
	}
}

func SetLogger(obj LogIF) {
	logger = obj
}

func Log(err error) {
	if logger == nil {
		return
	}
	logger.write(err)
}

func Print(err error) {
	fmt.Println(err)
}
