package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jangozw/gin-api-common/libs"
	"github.com/jangozw/gin-api-common/routes"
)

func main() {

	libs.Logger.Info("------------ starting server! --------------")
	engine := gin.New()
	routes.RegisterRouters(engine)
	if port, err := libs.Config.GetHttpPort(); err != nil {
		panic(err)
	} else if err := engine.Run(fmt.Sprintf(":%d", port)); err != nil {
		panic(err)
	}
	// Listening and serving HTTP on 127.0.0.1:{port}
}
