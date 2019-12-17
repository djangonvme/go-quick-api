package utils

import (
	//"net/http"
	//"a/c"
	"gin-api-common/consts"
	"time"
)

type ApiResponse struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data"`
}

func SuccessResponse(d interface{}) ApiResponse {
	return ApiResponse{
		consts.ApiCodeSuccess,
		consts.ApiCodeMaps[consts.ApiCodeSuccess],
		time.Now().Unix(),
		d,
	}
}

func SuccessResponseSimple() ApiResponse {
	return ApiResponse{
		consts.ApiCodeSuccess,
		consts.ApiCodeMaps[consts.ApiCodeSuccess],
		time.Now().Unix(),
		struct{}{},
	}
}

func FailResponse(errMsg string) ApiResponse {
	return FailResponseWithCode(consts.ApiCodeError, errMsg)
}

func FailResponseWithCode(c int, errMsg string) ApiResponse {
	var msg string
	if v, ok := consts.ApiCodeMaps[c]; ok {
		msg = v
	} else {
		msg = "api code unknown error"
	}
	return ApiResponse{
		c,
		msg + errMsg,
		time.Now().Unix(),
		struct{}{},
	}
}
