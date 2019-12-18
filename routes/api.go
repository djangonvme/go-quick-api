package routes

import (
	"gin-api-common/apis/v1"
	"gin-api-common/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRouters(router *gin.Engine) *gin.Engine {
	registerNoLogin(router)
	registerV1(router)
	return router
}

func registerV1(router *gin.Engine) {
	router.Group("/v1", middlewares.ApiMiddleware).
		POST("/logout", v1.Logout).
		GET("/user/list", v1.UserList).
		GET("/user/detail", v1.UserDetail)
}

func registerNoLogin(router *gin.Engine) {
	router.POST("/login", v1.Login)
}
