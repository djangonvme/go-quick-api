package middlewares

import (
	"gin-api-common/logs"
	"github.com/gin-gonic/gin"
	"time"
)

func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		//First we get info from api ctx, then let ctx go to it's deadline
		status := c.Writer.Status()
		ip := c.ClientIP()
		method := c.Request.Method
		uri := c.Request.URL.RequestURI()
		// 开始时间
		start := time.Now()
		// 处理请求
		c.Next()

		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)
		// 写日志的时间单独一个goroutine 处理，不占用接口调用时间
		logDone := make(chan bool)
		go func() {

			logs.Logger().Infof("| %3d | %3v | %5s | %s  %s |",
				status,
				latency,
				ip,
				method,
				uri,
			)
			logDone <- true
		}()
		select {
		case <-logDone:
			return
		case <-time.After(3 * time.Second): //3秒超时
			return
		}
	}
}
