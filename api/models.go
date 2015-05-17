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
