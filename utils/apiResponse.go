package utils

import (
	"gin-api-common/consts"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// api result to client format here

type ApiResponse struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data"`
}

type GinCtx struct {
	*gin.Context
}

//output data need use *gin.Context
// usage:
// utils.Response(c).Success("Ok")

func Response(c *gin.Context) *GinCtx {
	return &GinCtx{c}
}

func (c *GinCtx) Success(data interface{}) {
	c.JSON(http.StatusOK, ResponseSuccess(data))
}

func (c *GinCtx) SuccessSimple() {
	c.JSON(http.StatusOK, ResponseSuccessSimple())
}

func (c *GinCtx) Fail(errMsg interface{}) {
	c.FailWithCode(consts.ApiCodeError, errMsg)
}

func (c *GinCtx) FailWithCode(code int, errMsg interface{}) {
	c.JSON(http.StatusOK, ResponseFailWithCode(code, errMsg))
}

func ResponseSuccess(data interface{}) ApiResponse {
	msg, _ := consts.GetApiMsgByCode(consts.ApiCodeSuccess)
	return ApiResponse{
		consts.ApiCodeSuccess,
		msg,
		time.Now().Unix(),
		data,
	}
}

func ResponseSuccessSimple() ApiResponse {
	msg, _ := consts.GetApiMsgByCode(consts.ApiCodeSuccess)
	return ApiResponse{
		consts.ApiCodeSuccess,
		msg,
		time.Now().Unix(),
		struct{}{},
	}
}

func ResponseFailWithCode(code int, err interface{}) ApiResponse {
	msg, _ := consts.GetApiMsgByCode(code)
	return ApiResponse{
		code,
		msg + parseErrToMsg(err),
		time.Now().Unix(),
		struct{}{},
	}
}

func parseErrToMsg(err interface{}) (errMsg string) {
	if e, ok := err.(error); ok {
		errMsg = e.Error()
	} else if e, ok := err.(string); ok {
		errMsg = e
	} else {
		errMsg = ToJson(err)
	}
	return
}
