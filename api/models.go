package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// the Build struct that we can serialize and deserialize into Mongodb
type Build struct {
	ID          bson.ObjectId `form:"-" json:"id" bson:"_id,omitempty"`
	ProjectName string        `form:"project_name" json:"project_name" bson:"project_name"`
	Branch      string        `form:"branch" json:"branch" bson:"branch"`
	CommitSHA   string        `form:"commit_sha" json:"commit_sha" bson:"commit_sha"`
	Gcno        string        `form:"gcno" json:"gcno" bson:"gcno"`
}

type Runtime struct {
	Platform string
	Device   string
	SDK      string
}

type CoverageFileStat struct {
	Source        string
	lines         int
	lines_hit     int
	functions     int
	functions_hit int
}

type CoverageFunctionStat struct {
	Source  string
	name    string
	line_no int
	hit     int
}

type CoverageLineStat struct {
	Source  string
	line_no int
	hit     int
}

type CoverageTree struct {
	Files    []CoverageFileStat     `bson:,"inline"`
	Funtions []CoverageFunctionStat `bson:,"inline"`
	Lines    []CoverageLineStat     `bson:,"inline"`
}

type Report struct {
	Id          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	BuildId     bson.ObjectId
	RunAt       Runtime `bson:,"inline"`
	GcdaArchive string
}

func AllBuilds(db *mgo.Database) (builds []Build, err error) {
	err = db.C("builds").Find(nil).All(&builds)
	return builds, err
}

func FetchBuild(db *mgo.Database, buildId string) (build Build, err error) {
	err = db.C("builds").Find(bson.M{"_id": bson.ObjectIdHex(buildId)}).One(&build)
	return build, err
}

func InsertBuild(db *mgo.Database, build Build) (buildId bson.ObjectId, err error) {
	build.ID = bson.NewObjectId()
	err = db.C("builds").Insert(build)
	return build.ID, err
}
