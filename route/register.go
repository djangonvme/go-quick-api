package route

import (
	// "github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
	apisample "github.com/jangozw/go-quick-api/api/sample"
	"github.com/jangozw/go-quick-api/config"
	"github.com/jangozw/go-quick-api/middleware"
	"github.com/jangozw/go-quick-api/pkg/app"
)

// 注册路由, 先统一定义路由组，再给各个组定义具体路由，这样层次清晰
func Register(router *app.Engine) {
	var (
		// 公共中间件
		pubMiddleware = []gin.HandlerFunc{
			middleware.Header,
			middleware.LogRequest,
		}
		// /sample 路由
		sample = router.Group("/sample", pubMiddleware...)
		// /sample 需要登陆
		sampleLogin = router.Group("/sample", append(pubMiddleware, middleware.NeedLogin)...)
	)
	// /sample
	// curl http://127.0.0.1:8080/sample
	sample.GET("", func(c *gin.Context) (data interface{}, err error) {
		return "/sample welcome", nil
	})
	sample.GET("/status", func(c *gin.Context) (data interface{}, err error) {
		return config.GetAllStates(), nil
	})

	sample.POST("/login", apisample.Login)
	sampleLogin.POST("logout", apisample.Logout)
	sampleLogin.POST("/user", apisample.AddUser)
	sampleLogin.GET("/user/list", apisample.UserList)
	sampleLogin.GET("/user/detail", apisample.UserDetail)

	// open chrome http://127.0.0.1:8180/debug/pprof/
	// ginpprof.WrapGroup(router.Group("/debug").GinRouterGroup())
}
