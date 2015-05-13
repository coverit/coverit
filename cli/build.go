package main

import (
	"fmt"

	"github.com/codegangsta/cli"
)

func NewBuildCommand() cli.Command {

	return cli.Command{
		Name:  "build",
		Usage: "create a build",
		Action: func(c *cli.Context) {

			branchName := GetBranchName()
			commitSha := GetCommitSha()
			repoName := GetRepoName()

			fmt.Println("build " + repoName + " on " + branchName + " with commit " + commitSha)
		},
	}
}
