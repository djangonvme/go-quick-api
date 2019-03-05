package routes

import (
	"github.com/gin-gonic/gin"
	// 前面加. 表示直接可以用包里的函数名，不用package.IndexAPi()这么写
	"github.com/jangozw/gintest/middlewares"
	"github.com/jangozw/gintest/apis/v1/user"
)

func InitApiRouter(router *gin.Engine) *gin.Engine {
	//router := gin.Default()
	v1 := router.Group("/api/v1")
	v1.Use(middlewares.ApiMiddleware)
	v1.GET("/user/list", userApi.List)
	v1.GET("/user/detail", userApi.Detail)
	return router
}


