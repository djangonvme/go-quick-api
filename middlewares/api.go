package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/jangozw/gintest/apis"
)

func ApiMiddleware(c *gin.Context)  {
	token := c.Query("token")
	//中间件验证token
	if token == "" {
		c.AbortWithStatusJSON(http.StatusOK, apis.ErrorFormat(apis.CODE_TOKEN_VALID, "api token is required"))
		return
	}
	//继续下一步
	c.Next()
}



