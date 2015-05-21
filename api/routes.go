package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
)

func Include(m *martini.ClassicMartini) {
	m.Get("/builds", func(r render.Render, db *mgo.Database) {
		var builds, err = AllBuilds(db)
		if err == nil {
			r.JSON(200, builds)
		} else {
			r.JSON(400, nil)
		}
	})

	m.Post("/builds", binding.Form(Build{}), func(build Build, r render.Render, db *mgo.Database) {
		var buildId, err = InsertBuild(db, build)
		if err == nil {
			var build, _ = FetchBuild(db, buildId.Hex())
			r.JSON(200, build)
		} else {
			r.JSON(400, nil)
		}
	})

	m.Get("/builds/:build_id", func(params martini.Params, r render.Render, db *mgo.Database) {
		var build, err = FetchBuild(db, params["build_id"])
		if err == nil {
			r.JSON(200, build)
		} else {
			r.JSON(400, nil)
		}
	})

}
