package route

import (
	"context"
	"github.com/gin-gonic/gin"
	apiv1 "gitlab.com/task-dispatcher/api/v1"
	"gitlab.com/task-dispatcher/middleware"
	"gitlab.com/task-dispatcher/pkg/app"
)

func Register(ctx context.Context) func(router *app.Engine) {

	var (
		pubMiddleware = []gin.HandlerFunc{
			middleware.FixRequestBody(ctx),
			middleware.Header,
			middleware.LogRequest,
		}
	)
	return func(router *app.Engine) {
		var v1 = router.Group("/api/v1", append(pubMiddleware, middleware.CheckTaskRequest(ctx))...)

		v1.POST("/task/create", apiv1.TaskCreate)
		v1.GET("/task/result", apiv1.TaskResult)
		v1.POST("/task/apply", apiv1.TaskApply)
		v1.POST("/task/submit", apiv1.TaskSubmit)
	}

}
