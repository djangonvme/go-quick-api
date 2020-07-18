package middleware

import (
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jangozw/go-quick-api/pkg/app"
	"github.com/sirupsen/logrus"
)

func LogRequest(c *gin.Context) {
	// 处理请求
	c.Next()
	// 写日志的时间单独一个goroutine 处理，不占用接口调用时间
	apiLog(c)
}

// api 接口日志记录请求和返回
func apiLog(c *gin.Context) {
	user, _ := app.GetLoginUser(c)
	start := c.GetTime(app.CtxStartTime)
	// 执行时间
	latency := time.Now().Sub(start)
	resp, ok := c.Get(app.CtxKeyResponse)
	if !ok {
		resp = struct{}{}
	}
	var query interface{}
	if c.Request.Method == "GET" {
		query = c.Request.URL.Query()
	} else {
		postData, _ := c.GetRawData()
		query := make(map[string]interface{})
		json.Unmarshal(postData, &query)
	}

	// log 里有json.Marshal() 导致url转义字符
	app.Logger.WithFields(logrus.Fields{
		"uid":      user.ID,
		"query":    query,
		"response": resp,
	}).Infof("%s | %s |t=%3v | %s", c.Request.Method, c.Request.URL.RequestURI(), latency, c.ClientIP())
}
