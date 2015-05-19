package main

import (
	// "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type GitInfo struct {
	Branch string
	Commit string
	Author string
	Date   int64
}

// the Build struct that we can serialize and deserialize into Mongodb
type Build struct {
	Id          bson.ObjectId `form:"-" json:"id" bson:"_id,omitempty"`
	ProjectName string        `bson:"project_name"`
	Git         GitInfo
	Service     string
	Coverage    int
	Reports     []bson.ObjectId
}
