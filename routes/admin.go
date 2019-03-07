package routes

import (
	"gin-api-common/apis/admin/v1/user"
	"gin-api-common/middlewares"
	"github.com/gin-gonic/gin"
)

func InitAdminRouter(router *gin.Engine) *gin.Engine {
	v1 := router.Group("/admin/v1")

	v1.Use(middlewares.AdminMiddleware)

	v1.GET("/user/list", adminUserApi.List)

	//v1.POST("/helper/model")
	return router
}


