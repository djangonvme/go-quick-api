package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jangozw/gintest/apis/v1/user"
	// 前面加. 表示直接可以用包里的函数名，不用package.IndexAPi()这么写
)

func InitNoTokenRouter(router *gin.Engine) *gin.Engine {
	v1 := router.Group("/common/v1")
	v1.GET("/user/list", userApi.List)
	v1.GET("/user/detail", userApi.Detail)
	return router
}
