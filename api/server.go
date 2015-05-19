package main

import (
	"os"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
	// nothing here to configure
	}))

	session, err := mgo.Dial(os.Getenv("MONGO_URL")) // mongodb://localhost
	if err != nil {
		panic(err)
	}

	db := session.DB(os.Getenv("MONGO_DB"))

	buildService := BuildService{
		Resource: &BuildResource{db: db},
	}

	m.Group("/builds", func(r martini.Router) {
		r.Get("", buildService.Resource.ListBuilds)
		r.Post("", binding.Json(BuildCreateJson{}), buildService.CreateBuild)
	})

	m.Run()
}
