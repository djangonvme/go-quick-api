package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jangozw/gin-api-common/libs"
	"github.com/jangozw/gin-api-common/routes"
)

var (
	BuildVersion string //编译的app版本(makefile中自定义)
	BuildAt      string //编译时间
)

func main() {
	buildInfo()
	engine := gin.New()
	routes.RegisterRouters(engine)
	if port, err := libs.Config.GetHttpPort(); err != nil {
		log.Fatalln(err.Error())
	} else if err := engine.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalln(err.Error())
	}
	// Listening and serving HTTP on 127.0.0.1:{port}
}

func buildInfo() {
	libs.Logger.Info("app is starting", "BuildVersion=", BuildVersion, "BuildAt=", BuildAt)
}
