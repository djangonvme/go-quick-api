package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/jangozw/gin-api-common/consts"
	"github.com/jangozw/gin-api-common/services"
	"github.com/jangozw/gin-api-common/utils"
	"net/http"
)

func ApiMiddleware(c *gin.Context) {
	token := c.GetHeader(consts.HeaderKeyToken)
	if token == "" {
		c.AbortWithStatusJSON(http.StatusOK, utils.ResponseFailWithCode(consts.ApiCodeTokenValid, ""))
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
