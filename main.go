package main

import (
	"fmt"
	"gin-api-common/configs"
	"gin-api-common/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	routes.RegisterRouters(engine)
	if port, err := configs.GetHttpPort(); err != nil {
		panic(err)
	} else if err := engine.Run(fmt.Sprintf(":%d", port)); err != nil {
		panic(err)
	}
	// Listening and serving HTTP on 127.0.0.1:{port}
}
