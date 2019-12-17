package middlewares

import (
	"gin-api-common/consts"
	"gin-api-common/services"
	"gin-api-common/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusOK, utils.FailResponseWithCode(consts.ApiCodeTokenValid, "token is required"))
		return
	}
	//check login user
	if jwt, err := services.VerifyAppToken(token); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, utils.FailResponseWithCode(consts.ApiCodeTokenValid, err.Error()))
		return
	} else {
		c.Set("login_user", jwt.UserId)
		//继续下一步
		c.Next()
	}
}
