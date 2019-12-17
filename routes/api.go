package routes

import (
	"gin-api-common/apis"
	"gin-api-common/middlewares"
	"github.com/gin-gonic/gin"
)

func InitApiRouters(router *gin.Engine) *gin.Engine {
	router.POST("/v1/login", apis.Login)
	v1 := router.Group("/v1")
	v1.Use(middlewares.ApiMiddleware)
	v1.POST("/logout", apis.Logout)
	v1.GET("/user/list", apis.UserList)
	return router
}
