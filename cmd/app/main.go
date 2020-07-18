package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/jangozw/go-quick-api/param"
	"github.com/jangozw/go-quick-api/pkg/app"
	"github.com/jangozw/go-quick-api/pkg/util"
	"github.com/jangozw/go-quick-api/route"
	"github.com/urfave/cli/v2"
)

var (
	// 编译的app版本
	BuildVersion string
	// 编译时间
	BuildAt string
)

func main() {
	if err := stack().Run(os.Args); err != nil {
		panic(err)
	}
}

func stack() *cli.App {
	app.BuildInfo = fmt.Sprintf("%s-%s-%s-%s-%s", runtime.GOOS, runtime.GOARCH, BuildVersion, BuildAt, util.Now())
	return &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        param.ArgConfig,
				Value:       param.ArgConfigFilename,
				Destination: &app.ConfigFile,
			},
		},
		Name:    param.AppName,
		Version: app.BuildInfo,
		Usage:   fmt.Sprintf("./%s -%s=%s", param.AppName, param.ArgConfig, param.ArgConfigFilename),
		Action: func(c *cli.Context) error {
			// 初始化服务，注册路由
			app.Init()
			return app.NewGin(route.Register).Run()
		},
	}
}
