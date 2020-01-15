package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jangozw/gin-api-common/libs"
	"github.com/jangozw/gin-api-common/routes"
	"log"
)

// build info
var (
	Version string
	Build   string
)

func main() {
	libs.Logger.Info("Starting server...", "version=", Version, "build=", Build)
	engine := gin.New()
	routes.RegisterRouters(engine)
	if port, err := libs.Config.GetHttpPort(); err != nil {
		log.Fatalln(err.Error())
	} else if err := engine.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalln(err.Error())
	}
	// Listening and serving HTTP on 127.0.0.1:{port}
}
