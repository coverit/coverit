package main

import (
	"fmt"

  "github.com/codegangsta/cli"
)


func NewReportCommand() cli.Command {

  return cli.Command{

    Name:  "report",
		Usage: "create a report of code coverage",
		Action: func(c *cli.Context) {

      fmt.Println("Hi new report")
    },
  }

}
