package models

import (
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
)


// the Build struct that we can serialize and deserialize into Mongodb
type Build struct {
  ID      bson.ObjectId `form:"-" json:"id" bson:"_id,omitempty"`
  ProjectName string `form:"project_name" json:"project_name" bson:"project_name"`
  Branch  string `form:"branch" json:"branch" bson:"branch"`
  CommitSHA string `form:"commit_sha" json:"commit_sha" bson:"commit_sha"`
  Gcno string `form:"gcno" json:"gcno" bson:"gcno"`
}

func AllBuilds(db *mgo.Database) []Build {
  var builds []Build
  db.C("builds").Find(nil).All(&builds)
  return builds
}

func FetchBuild(db *mgo.Database, build_id string) Build {
  var build Build
  db.C("builds").Find(bson.M{"_id": build_id}).One(&build)
  return build
}

func InsertBuild(db *mgo.Database, build Build) {
  db.C("builds").Insert(build)
}
