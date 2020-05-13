package routes

import (
	"github.com/gin-gonic/gin"
	v0 "github.com/jangozw/gin-api-common/apis/v0"
	"github.com/jangozw/gin-api-common/middlewares"
	"github.com/jangozw/gin-api-common/utils"
	"time"
)

func RegisterRouters(router *gin.Engine) *gin.Engine {
	router.Use(middlewares.CommonMiddleware, middlewares.LoggerToFile())
	registerUnLogin(router)
	registerV0(router)
	return router
}
// 需要登陆
func registerV0(router *gin.Engine) {
	router.Group("/v0", middlewares.CheckJwtLogin).
		POST("/logout", utils.WithAPIHandlerFunc(v0.Logout)).
		GET("/user/list", utils.WithAPIHandlerFunc(v0.UserList)).
		GET("/user/detail", utils.WithAPIHandlerFunc(v0.UserDetail)).
		POST("/user/add", utils.WithAPIHandlerFunc(v0.AddUser))
}

// 无需登陆
func registerUnLogin(router *gin.Engine) {
	router.GET("/", utils.WithAPIHandlerFunc(func(c *utils.ApiContext) {
		c.Success(map[string]string{
			"title":   "Welcome! ",
			"time":    time.Now().Format(utils.YMDHIS),
			"buildVersion": utils.Build.Version,
			"buildAt": utils.Build.Time,
		})
	}))

	router.Group("/v0").	POST("/login", utils.WithAPIHandlerFunc(v0.Login)).
		GET("/timeout", utils.WithAPIHandlerFunc(v0.TimeOutOperation))
}

