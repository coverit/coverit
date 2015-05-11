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

			branchName := GetBranchName()
			commitSha := GetCommitSha()
			repoName := GetRepoName()

			createReport(branchName, repoName, commitSha)
	  },
  }

}

func createReport(branch string, repo string, commit string) error {

	fmt.Println("report " + repo + " on " + branch + " with commit " + commit)

	return nil
}
