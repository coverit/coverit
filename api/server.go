package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func init() {
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
	// nothing here to configure
	}))
	// use the Mongo middleware
	m.Use(DB())

	Include(m)

	m.Run()
}
