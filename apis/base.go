package apis

import (
	"gin-api-common/consts"
	"github.com/gin-gonic/gin"
)

//
func getLoginUid(c *gin.Context) int64 {
	return c.GetInt64(consts.CtxKeyLoginUser)
}
