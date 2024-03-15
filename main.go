package main

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"gitlab.com/qubic-pool/config"
	"gitlab.com/qubic-pool/pkg/app"
	"gitlab.com/qubic-pool/pkg/db"
	"gitlab.com/qubic-pool/pkg/logger"
	"gitlab.com/qubic-pool/pkg/redis"
	"gitlab.com/qubic-pool/route"
	"log"
	"os"
)

var (
	BuildVersion string
	BuildAt      string
)

const (
	InitRedis = iota //0
	InitDb           //1
)

func main() {
	ins := cli.NewApp()
	ins.Name = "qubic-pool"
	ins.Version = BuildVersion + "," + BuildAt
	ins.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "config",
			Aliases:     []string{"c"},
			Value:       "config.toml",
			Required:    true,
			DefaultText: "-c=config.toml",
		},
	}
	ins.Action = func(c *cli.Context) error {
		return start(c,
			InitDb,
			InitRedis,
		)
	}

	// tasks
	{
		//go service.ProcessWaitingSubmitSolutions()
	}

	// start server
	err := ins.Run(os.Args)
	if err != nil {
		log.Fatalf("%v\n", err.Error())
	}
}
func start(c *cli.Context, inits ...int) error {
	// load config
	cfg, err := config.LoadConfig(c.String("c"))
	if err != nil {
		return errors.Wrap(err, "LoadConfig failed!")
	}
	// init logger
	if err := logger.InitLogger(cfg.Log.Dir); err != nil {
		return errors.Wrap(err, "InitLogger failed!")
	}
	b, _ := json.Marshal(cfg)
	logger.Instance.Infof("load config: %s", b)

	for _, v := range inits {
		if v == InitDb {
			// init db. open it if you config mysql
			if err := db.InitDB(cfg.MySQL.Host, cfg.MySQL.User, cfg.MySQL.Password, cfg.MySQL.DbName, logger.Instance); err != nil {
				return errors.Wrap(err, "InitDb failed!")
			}
		}
		if v == InitRedis {
			// init redis. open it if you config redis
			if err := redis.InitRedis(cfg.Redis.Host, cfg.Redis.Password, cfg.Redis.Db, cfg.Redis.PoolSize); err != nil {
				return errors.Wrap(err, "InitRedis failed")
			}
		}
	}
	return app.NewGin(route.Register(c.Context)).Run(config.Instance.Server.Listen)
}
