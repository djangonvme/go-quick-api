package utils

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jangozw/gin-api-common/consts"
	"net/http"
	"time"
)

// 请求上下文, 基于 gin.Context
type ApiContext struct {
	*gin.Context
}

// 各个接口函数类型
type APIHandlerFunc func(c *ApiContext)

// 带有请求超时控制的自定义的api handler
func APIHandler(handler APIHandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		handler(&ApiContext{c})
	}
}

// api result to client format here
type ApiResponse struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data"`
}

func (c *ApiContext) Success(data interface{}) {
	result := ResponseSuccess(data)
	c.OutputJSON(http.StatusOK, result)

}
func (c *ApiContext) SuccessSimple() {
	c.OutputJSON(http.StatusOK, ResponseSuccessSimple())
}
func (c *ApiContext) Fail(errMsg interface{}) {
	c.FailWithCode(consts.ApiCodeError, errMsg)
}

func (c *ApiContext) FailWithCode(code int, errMsg interface{}) {
	c.OutputJSON(http.StatusOK, ResponseFailWithCode(code, errMsg))
}

func (c *ApiContext) OutputJSON(code int, obj interface{}) {
	c.JSON(code, obj)
	// 请求超时,已经在中间件输出了，这里不再处理
	if c.Request.Context().Err() != context.DeadlineExceeded {
		//c.JSON(code, obj)
	}
}

//当前登陆的用户id
func (c *ApiContext) GetLoginUid() int64 {
	return c.GetInt64(consts.CtxKeyLoginUser)
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
		msg + " " + parseErrToMsg(err),
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
