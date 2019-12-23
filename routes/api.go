package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jangozw/gin-api-common/apis/v1"
	"github.com/jangozw/gin-api-common/middlewares"
	"net/http"
)

func RegisterRouters(router *gin.Engine) *gin.Engine {
	registerNoLogin(router)
	registerV1(router)
	return router
}

func registerV1(router *gin.Engine) {
	router.Group("/v1", middlewares.LoggerToFile(), middlewares.ApiMiddleware).
		POST("/logout", v1.Logout).
		GET("/user/list", v1.UserList).
		GET("/user/detail", v1.UserDetail)
}

func registerNoLogin(router *gin.Engine) {
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "welcome!"})
	})
	router.POST("/user/add", v1.AddUser)
	router.POST("/login", v1.Login)
}
