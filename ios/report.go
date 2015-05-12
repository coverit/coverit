package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
)

func NewReportCommand() cli.Command {

	return cli.Command{

		Name:  "report",
		Usage: "create a report of code coverage",
		Action: func(c *cli.Context) {

			branch := GetBranchName()
			commit := GetCommitSha()
			repo := GetRepoName()

			if branch == "" {
				log.Fatal("Please specify a correct branch name")
			}

			if repo == "" {
				log.Fatal("Please specify a correct repo name")
			}

			if commit == "" {
				log.Fatal("Please specify a correct commit sha")
			}

			createReport(branch, repo, commit)
		},
	}

}

func createReport(branch string, repo string, commit string) error {

	fmt.Println("reporting " + repo + " on " + branch + " with commit " + commit)

	var lcovPath = "./lcov.zip" //Fixme: get real lcov path

	extraParams := map[string]string{
		"project_name": repo,
		"commit_sha":   commit,
		"branch":       branch,
	}

	request, err := newfileUploadRequest("http://google.com/upload", extraParams, "lcov", lcovPath)

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(request)

	if err != nil {
		log.Fatal(err)
	}

	body := &bytes.Buffer{}
	_, err = body.ReadFrom(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	resp.Body.Close()
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header)
	fmt.Println(body)

	return nil
}

func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return http.NewRequest("POST", uri, body)
}
