package main

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type CovYml struct {
	Repo   string `yaml:"coverit_repo_name"`
	Commit string `yaml:"coverit_commit_sha"`
	Branch string `yaml:"coverit_branch_name"`
	Token  string `yaml:"coverit_api_token"`
}

func GetBranchName() string {

	branchName := os.Getenv("coverit_branch_name")

	if branchName != "" {
		return branchName
	}

	cov, _ := coveritYml()

	return cov.Branch
}

func GetCommitSha() string {
	commitSha := os.Getenv("coverit_commit_sha")

	if commitSha != "" {
		return commitSha
	}

	cov, _ := coveritYml()

	return cov.Commit
}

func GetRepoName() string {
	repoName := os.Getenv("coverit_repo_name")

	if repoName != "" {
		return repoName
	}

	cov, _ := coveritYml()

	return cov.Repo
}

func coveritYml() (CovYml, error) {

	cov := CovYml{}
	fileName := ".coverit.yml"

	source, err := ioutil.ReadFile(fileName)

	if err != nil {
		return cov, err
	}

	yaml.Unmarshal(source, &cov)

	return cov, nil
}
