package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jangozw/gin-api-common/libs"
	"github.com/jangozw/gin-api-common/routes"
	"github.com/jangozw/gin-api-common/utils"
)

var (
	BuildVersion string //编译的app版本
	BuildAt      string //编译时间
)

func init() {
	// 配置文件路径，取命令行config参数作为路径
	cmdArgsConfig := flag.String("config", libs.ConfigFile, "config file path, default: "+libs.ConfigFile)
	flag.Parse()
	if cmdArgsConfig != nil {
		libs.ConfigFile = *cmdArgsConfig
	}
	// 构建信息
	utils.SetBuildInfo(BuildVersion, BuildAt)
}

func main() {
	// http serve
	engine := gin.New()
	routes.RegisterRouters(engine)
	if port, err := libs.Config.GetHttpPort(); err != nil {
		panic(err)
	} else if err := engine.Run(fmt.Sprintf(":%d", port)); err != nil {
		panic(err)
	}
}
