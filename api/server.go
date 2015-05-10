package main

import (
	"fmt"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/tmilewski/goenv"

	"github.com/coverit/coverit/api/rest"
)

func init() {
	err := goenv.Load()

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
	// nothing here to configure
	}))
	// use the Mongo middleware
	m.Use(rest.DB())

	rest.Include(m)

	m.Run()
}
