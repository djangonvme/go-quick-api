package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jangozw/gin-api-common/utils"
)

func CommonMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Header("build-version", utils.Build.Version)
	c.Header("build-at", utils.Build.Time)
	c.Next()
}
