package middleware

import (
	"github.com/jangozw/go-quick-api/erron"
	"github.com/jangozw/go-quick-api/param"

	"github.com/gin-gonic/gin"

	"github.com/jangozw/go-quick-api/pkg/app"
)

func NeedLogin(c *gin.Context) {
	token := c.GetHeader(param.TokenHeaderKey)
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
