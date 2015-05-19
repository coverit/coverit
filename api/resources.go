package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type BuildResource struct {
	db *mgo.Database
}

func (res *BuildResource) CreateBuild(project string, git GitInfo, service string) (Build, error) {

	build := Build{}
	build.ProjectName = project
	build.Git = git
	build.Service = service

	build.Id = bson.NewObjectId()

	err := res.db.C("builds").Insert(build)
	return build, err
}

func (res *BuildResource) ListBuilds(project string) ([]Build, error) {

	var builds []Build

	err := res.db.C("builds").Find(bson.M{"project_name": project}).All(&builds)
	return builds, err
}
