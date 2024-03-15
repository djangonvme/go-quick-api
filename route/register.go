package route

import (
	"context"
	apiv1 "gitlab.com/qubic-pool/api/v1"
	"gitlab.com/qubic-pool/middleware"
	"gitlab.com/qubic-pool/pkg/app"
)

func Register(ctx context.Context) func(router *app.Engine) {
	return func(router *app.Engine) {
		var v1 = router.Group("/api/v1",
			middleware.SetHeader,
		)

		v1.POST("/user/register", apiv1.UserRegister)
		v1.POST("/user/login", apiv1.UserLogin)

		var v1Login = router.Group("/api/v1",
			middleware.Common,
			middleware.SetHeader,
			middleware.CheckLogin,
		)

		v1Login.GET("/user/info", apiv1.GetUserInfo)
		v1Login.POST("/user/update", apiv1.UserUpdateInfo)

	}

}
