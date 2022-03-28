package middleware

import (
	"github.com/go-quick-api/erron"
	"github.com/go-quick-api/types"

	"github.com/gin-gonic/gin"

	"github.com/go-quick-api/pkg/app"
)

func NeedLogin(c *gin.Context) {
	token := c.GetHeader(types.TokenHeaderKey)
	if token == "" {
		app.AbortJSON(c, app.ResponseFailByCode(erron.UnLogin))
		return
	}
	if _, err := app.ParseUserByToken(token); err != nil {
		app.AbortJSON(c, app.ResponseFail(erron.New(erron.UnLogin, err.Error())))
		return
	}
	// 继续下一步
	c.Next()
}
