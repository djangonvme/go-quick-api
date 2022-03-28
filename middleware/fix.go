package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-quick-api/pkg/app"
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
