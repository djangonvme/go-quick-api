package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-quick-api/pkg/app"
)

func LogRequest(c *gin.Context) {
	// 处理请求
	c.Next()
	// 写日志的时间单独一个goroutine 处理，不占用接口调用时间
	app.ApiLog(c)
}
