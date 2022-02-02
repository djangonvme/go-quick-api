package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"gitlab.com/task-dispatcher/pkg/app"
	"log"
)

func FixRequestBody(ctx context.Context) func(c *gin.Context) {
	return func(c *gin.Context) {
		_, err := app.GetSaveRawData(c)
		if err != nil {
			log.Printf("GetSaveRawData failed: %v", err)
		}
	}
}
