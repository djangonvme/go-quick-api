package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"gitlab.com/task-dispatcher/cmd/clis"
	"gitlab.com/task-dispatcher/pkg/app"
	"log"
	"os"
)

var (
	BuildVersion string
	BuildAt      string
)

func main() {
	app.BuildInfo = fmt.Sprintf("%v,%v", BuildVersion, BuildAt)
	c := cli.NewApp()
	c.Name = "task-dispatcher"
	c.Version = BuildVersion + "," + BuildAt
	c.Usage = ""
	c.Commands = []*cli.Command{
		clis.RunCmd,
	}
	err := c.Run(os.Args)
	if err != nil {
		log.Fatalf("app run failed: %v\n", err.Error())
	}
}
