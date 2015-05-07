package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
)

var (
	version = "0.0.0"
)

func NewBuildCommand() cli.Command {

	return cli.Command{
		Name:  "build",
		Usage: "create a build",
		Action: func(c *cli.Context) {
			fmt.Println("hi")
			collectGcno()
		},
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "coverit"
	app.Version = version
	app.Usage = "command line utility"

	app.Commands = []cli.Command{
		NewBuildCommand(),
	}
	app.Run(os.Args)
}
