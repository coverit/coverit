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

	var mongoURL string
	if len(os.Getenv("MONGO_DB")) != 0 {
		mongoURL = os.Getenv("MONGO_DB")
	} else if len(os.Getenv("MONGO_PORT_27017_TCP_ADDR")) != 0 {
		mongoURL = os.Getenv("MONGO_PORT_27017_TCP_ADDR")
	} else {
		mongoURL = "localhost"
	}

	session, err := mgo.Dial(mongoURL) // mongodb://localhost
	if err != nil {
		panic(err)
	}

	db := session.DB(mongoURL)

	buildService := BuildService{
		Resource: &BuildResource{db: db},
	}

	m.Group("/builds", func(r martini.Router) {
		r.Get("", buildService.Resource.ListBuilds)
		r.Post("", binding.Json(BuildCreateJson{}), buildService.CreateBuild)
	})

	m.Run()
}
