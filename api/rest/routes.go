package rest

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"

	"./models"
)

func Include(m *martini.ClassicMartini) {
	m.Get("/builds", func(r render.Render, db *mgo.Database) {
		r.JSON(200, models.AllBuilds(db))
	})

	m.Post("/builds", binding.Form(models.Build{}), func(build models.Build, r render.Render, db *mgo.Database) {
		models.InsertBuild(db, build)
		r.JSON(200, models.AllBuilds(db))
	})

	m.Get("/builds/:build_id", func(params martini.Params, r render.Render, db *mgo.Database) {
		r.JSON(200, models.FetchBuild(db, params["build_id"]))
	})
}
