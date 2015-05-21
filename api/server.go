package main

import (
	"os"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
)

func Init() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
	// nothing here to configure
	}))
	// use the Mongo middleware
	// m.Use(DB())

	// Include(m)

	session, err := mgo.Dial(os.Getenv("MONGO_URL")) // mongodb://localhost
	if err != nil {
		panic(err)
	}

	db := session.DB(os.Getenv("MONGO_DB"))

	reportResource := &ReportResource{db: db}

	m.Group("/reports", func(r martini.Router) {
		r.Get("/:id", reportResource.GetReport)
		r.Post("/", reportResource.GetReport)
		r.Put("/:id", reportResource.GetReport)
		r.Delete("/:id", reportResource.GetReport)
	})

	m.Run()
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
	// nothing here to configure
	}))
	// use the Mongo middleware
	// m.Use(DB())

	// Include(m)

	session, err := mgo.Dial(os.Getenv("MONGO_URL")) // mongodb://localhost
	if err != nil {
		panic(err)
	}

	db := session.DB(os.Getenv("MONGO_DB"))

	reportResource := &ReportResource{db: db}

	m.Group("/reports", func(r martini.Router) {
		r.Get("/:id", reportResource.GetReport)
		r.Post("/", reportResource.CreateReport)
		r.Put("/:id", reportResource.GetReport)
		r.Delete("/:id", reportResource.GetReport)
	})

	m.Run()
}
