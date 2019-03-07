package routes

import (
	"gin-api-common/apis/v1/user"
	"github.com/gin-gonic/gin"
)

func InitNoTokenRouter(router *gin.Engine) *gin.Engine {
	v1 := router.Group("/common/v1")
	v1.GET("/user/list", userApi.List)
	v1.GET("/user/detail", userApi.Detail)
	return router
}
