package middlewares

import "github.com/gin-gonic/gin"

func CommonMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Next()
}
