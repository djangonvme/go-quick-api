package middlewares

import (
	"gin-api-common/consts"
	"gin-api-common/services"
	"gin-api-common/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ApiMiddleware(c *gin.Context) {
	token := c.GetHeader(consts.HeaderKeyToken)
	if token == "" {
		c.AbortWithStatusJSON(http.StatusOK, utils.ResponseFailWithCode(consts.ApiCodeTokenValid, "token is required"))
		return
	}
	//check login user
	if jwt, err := services.VerifyAppToken(token); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, utils.ResponseFailWithCode(consts.ApiCodeTokenValid, err))
		return
	} else {
		//verify token success, then set the login user for context
		c.Set(consts.CtxKeyLoginUser, jwt.UserId)
		//继续下一步
		c.Next()
	}
}
