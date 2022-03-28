package main

import (
	"fmt"
	"github.com/go-quick-api/cmd/clis"
	"github.com/go-quick-api/pkg/app"
	"github.com/urfave/cli/v2"
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
	c.Name = "go-quick-api"
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
