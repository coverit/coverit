package main

import (
	"fmt"
	"os"
  "io/ioutil"

	"github.com/codegangsta/cli"
  "gopkg.in/yaml.v2"
)

type CovYml struct {
  Repo   string `yaml:"COVERIT_REPO_NAME"`;
  Commit string `yaml:"COVERIT_COMMIT_SHA"`;
  Branch string `yaml:"COVERIT_BRANCH_NAME"`;
  Token  string `yaml:"COVERIT_API_TOKEN"`;
}

func NewBuildCommand() cli.Command {

	return cli.Command{
		Name:  "build",
		Usage: "create a build",
		Action: func(c *cli.Context) {

			branchName := getBranchName()
			commitSha := getCommitSha()
			repoName := getRepoName()

			fmt.Println("build " + repoName + " on " + branchName + " with commit " + commitSha)

      run()
    },
	}
}

func getBranchName() string {

  branchName := os.Getenv("COVERIT_BRANCH_NAME")

  if branchName != "" {
    return branchName
  }

  cov := coveritYml()

  return cov.Branch
}

func getCommitSha() string {
  commitSha := os.Getenv("COVERIT_COMMIT_SHA")

  if commitSha != "" {
    return commitSha
  }

  cov := coveritYml()

  return cov.Commit
}

func getRepoName() string {
  repoName := os.Getenv("COVERIT_REPO_NAME")

  if repoName != "" {
    return repoName
  }

  cov := coveritYml()

  return cov.Repo
}

func coveritYml() CovYml {

  cov := CovYml{}
  fileName := ".coverit.yml" // Fixme: How to define the file path?

  source, err := ioutil.ReadFile(fileName)

  if err != nil {
    panic(err)
  }

  yaml.Unmarshal(source, &cov)

  return cov
}

func run() {

}
