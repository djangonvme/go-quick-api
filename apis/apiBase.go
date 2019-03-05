package apis

import (
	//"net/http"
	//"a/c"
)

const CODE_SUCCESS = 200
const CODE_ERROR = 400
const CODE_TOKEN_VALID = 401

type ApiBase struct {
	Code int
	Msg string
	Data interface{}
}

func SuccessFormat (d interface{}) ApiBase{
	a := ApiBase{
		200,
		"请求成功",
		d,
	}
	return a
}
func ErrorFormat (c int, d interface{}) ApiBase{
	a := ApiBase{
		c,
		"请求失败",
		d,
	}
	return a
}
