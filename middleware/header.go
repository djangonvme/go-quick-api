package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-quick-api/erron"
	"github.com/go-quick-api/pkg/app"
)

func Header(c *gin.Context) {
	if c.Request.Method != "GET" && c.GetHeader("Content-Type") != "application/json" {
		app.AbortJSON(c, app.ResponseFail(erron.New(erron.ErrRequestParam, "invalid Content-Type")))
		return
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Header("Build-Version", app.BuildInfo)
	c.Next()
}
