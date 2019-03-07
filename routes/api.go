package routes

import (
	"gin-api-common/apis/v1/user"
	"gin-api-common/middlewares"
	"github.com/gin-gonic/gin"
)

func InitApiRouter(router *gin.Engine) *gin.Engine {
	//router := gin.Default()
	v1 := router.Group("/api/v1")
	v1.Use(middlewares.ApiMiddleware)
	v1.GET("/user/list", userApi.List)
	v1.GET("/user/detail", userApi.Detail)
	return router
}


