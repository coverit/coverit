package main

import (
	"os"

	"github.com/codegangsta/cli"
)

var (
	// commit sha for the current build.
	version string
	sha     string
)

func main() {
	app := cli.NewApp()
	app.Name = "coverit"
	app.Version = version + "@" + sha
	app.Usage = "command line utility"

	app.Commands = []cli.Command{
		NewBuildCommand(),
		NewReportCommand(),
	}
	app.Run(os.Args)
}
