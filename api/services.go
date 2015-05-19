package main

import (
	"fmt"
	"net/http"

	"github.com/martini-contrib/render"
)

type BuildService struct {
	Resource *BuildResource
}

type BuildCreateJson struct {
	ProjectName string `json:"project_name"`
	Git         GitInfo
	Service     string
}

func (service *BuildService) CreateBuild(body BuildCreateJson, r render.Render) {

	build, err := service.Resource.CreateBuild(body.ProjectName, body.Git, body.Service)

	if err != nil {
		r.JSON(400, nil)
	} else {
		r.JSON(201, build)
	}
}

func (service *BuildService) ListBuilds(req *http.Request, r render.Render) {

	projectName := req.URL.Query().Get("project_name")
	fmt.Println("project: " + projectName)
	if len(projectName) == 0 {
		r.JSON(400, nil)
		return
	}

	builds, err := service.Resource.ListBuilds(projectName)
	if err != nil {
		r.JSON(404, nil)
	} else {
		r.JSON(200, builds)
	}

}
