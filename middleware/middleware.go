package middleware

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/qubic-pool/erron"
	"gitlab.com/qubic-pool/pkg/app"
	"gitlab.com/qubic-pool/service"
	"net/http"
	"strings"
)

func CheckLogin(c *gin.Context) {

	token := c.GetHeader("Authorization")
	//logger.Instance.Infof("CheckLogin header token: %s", token)

	token = strings.ReplaceAll(strings.TrimPrefix(token, "Bearer"), " ", "")
	if user, err := service.CheckLoginUserByToken(token); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, app.ResponseFailByCode(erron.UnLogin))
		return
	} else {
		c.Set("loginUser", *user)
	}
	// 继续下一步
	c.Next()
}

func SetHeader(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.Next()
}

func Common(c *gin.Context) {
	app.GetRequestBody(c)
	c.Next()
}
