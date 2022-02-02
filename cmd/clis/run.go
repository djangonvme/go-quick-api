package clis

import (
	"context"
	"fmt"
	"github.com/urfave/cli/v2"
	"gitlab.com/task-dispatcher/config"
	"gitlab.com/task-dispatcher/pkg/app"
	"gitlab.com/task-dispatcher/pkg/singleton"
	"gitlab.com/task-dispatcher/route"
	"gitlab.com/task-dispatcher/service"
	"gitlab.com/task-dispatcher/types"
	"go.uber.org/fx"
	"log"
)

var RunCmd = &cli.Command{
	Name:  "run",
	Usage: "",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "config",
			Value:    "config.toml",
			Required: true,
			Usage:    "--config=/xx/xxx/config.toml",
		},
		&cli.BoolFlag{
			Name:     types.TaskTypeLotusCommit,
			Value:    false,
			Required: false,
			Usage:    "--lotus-commit2=true",
		},
	},
	Action: func(c *cli.Context) error {
		tasks := map[string]bool{
			types.TaskTypeLotusCommit: c.Bool(types.TaskTypeLotusCommit),
		}
		var has bool
		for _, v := range tasks {
			if v == true {
				has = true
				break
			}
		}
		if !has {
			return fmt.Errorf("no tasks configed for running! exit")
		}
		ctx := context.WithValue(c.Context, types.KeyAllowTasks, tasks)
		container := fx.New(
			// load or new global resource
			fx.Provide(config.LoadConfig(c.String("config"))),
			fx.Provide(singleton.NewLogger("task-dispatcher")),
			fx.Provide(singleton.NewDB),
			fx.Provide(singleton.NewRedis),

			// global resource instances
			fx.Populate(&app.CfgInstance),
			fx.Populate(&app.DBInstance),
			fx.Populate(&app.LoggerInstance),
			fx.Populate(&app.RedisInstance),

			// invoke serve
			fx.Invoke(register),
		)
		container.Start(ctx)
		defer container.Stop(ctx)
		return nil
	},
}

func register(lc fx.Lifecycle) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				for _, h := range getTaskHandlers(ctx) {
					go h.Revert()
				}
				return app.NewGin(route.Register(ctx)).Run()
			},
			OnStop: func(ctx context.Context) error {
				// return server.Shutdown(ctx)
				log.Println("application stopped")
				return nil
			},
		},
	)
}

func getTaskHandlers(ctx context.Context) []types.TaskManager {
	var hds []types.TaskManager
	val := ctx.Value(types.KeyAllowTasks)
	mp, _ := val.(map[string]bool)
	for k, v := range mp {
		if v {
			hd, err := service.GetTaskHandler(k)
			if err != nil {
				log.Printf("GetTaskHandler by taskType(%v) failed: %v\n", k, err)
				continue
			} else {
				hds = append(hds, hd)
			}
		}
	}
	return hds
}
