package main

import (
	"os"

	"github.com/codegangsta/cli"
)

var (
	version = "0.0.0"
)

func main() {
	app := cli.NewApp()
	app.Name = "coverit"
	app.Version = version
	app.Usage = "command line utility"

	app.Commands = []cli.Command{
		NewBuildCommand(),
		NewReportCommand(),
	}
	app.Run(os.Args)
}
