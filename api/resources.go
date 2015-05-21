package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ReportResource struct {
	db *mgo.Database
}

func (res *ReportResource) GetReport(params martini.Params, r render.Render) {

	var report Report

	err := res.db.C("reports").Find(bson.M{"_id": bson.ObjectIdHex(params["id"])}).One(&report)

	if err != nil {
		r.JSON(404, nil)
	} else {
		r.JSON(200, report)
	}

}

func (res *ReportResource) CreateReport(r render.Render) {

	report := Report{}
	report.Id = bson.NewObjectId()

	err := res.db.C("reports").Insert(report)

	if err != nil {
		r.JSON(500, nil)
	} else {
		r.JSON(201, report)
	}

}
