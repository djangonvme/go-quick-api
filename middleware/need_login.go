package middleware

import (
	"gitlab.com/task-dispatcher/erron"
	"gitlab.com/task-dispatcher/types"

	"github.com/gin-gonic/gin"

	"gitlab.com/task-dispatcher/pkg/app"
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
