package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jangozw/gintest/apis/admin/v1/user"
	"github.com/jangozw/gintest/middlewares"
)

func InitAdminRouter(router *gin.Engine) *gin.Engine {
	v1 := router.Group("/admin/v1")

	v1.Use(middlewares.AdminMiddleware)

	v1.GET("/user/list", adminUserApi.List)

	//v1.POST("/helper/model")
	return router
}


