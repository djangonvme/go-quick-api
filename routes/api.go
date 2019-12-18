package routes

import (
	"gin-api-common/apis"
	"gin-api-common/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRouters(router *gin.Engine) *gin.Engine {
	registerNoLogin(router)
	v1 := router.Group("/v1")
	v1.Use(middlewares.ApiMiddleware)
	registerV1(v1)
	return router
}

func registerV1(v1 *gin.RouterGroup) {
	v1.POST("/logout", apis.Logout)
	v1.GET("/user/list", apis.UserList)
	v1.GET("/user/detail", apis.UserDetail)
}

func registerNoLogin(router *gin.Engine) {
	router.POST("/login", apis.Login)
}
