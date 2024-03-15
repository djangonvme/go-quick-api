package erron

type ErrorIF interface {
	Error() string
	Code() ErrCode
}

type Error struct {
	code ErrCode
	msg  string
}

func (e *Error) Error() string {
	return e.msg
}

func (e *Error) Code() ErrCode {
	return e.code
}

func New(code ErrCode, msg string) *Error {
	return &Error{
		code: code,
		msg:  msg,
	}
}
