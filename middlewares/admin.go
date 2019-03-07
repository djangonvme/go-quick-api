package middlewares

import (
	"gin-api-common/apis"
	"gin-api-common/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminMiddleware(c *gin.Context)  {
	token := c.Query("token")
	t := models.Token{}.GetByToken(token)
	errStr := ""
	//中间件验证token
	if t.Id == 0 {
		errStr = "token is required or invalid"
	} else if t.IsTokenExpired() {
		errStr = "token has expired"
	}
	if errStr != "" {
		c.AbortWithStatusJSON(http.StatusOK, apis.ErrorFormat(apis.CODE_TOKEN_VALID, errStr))
		//return
	}
	//继续下一步
	c.Next()
}

