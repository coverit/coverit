package main

import (
  "net/http"
  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/render"

  "./rest"
)

func main() {

  m := martini.Classic()
  m.Use(render.Renderer(render.Options {
    // nothing here to configure
  }))
  // use the Mongo middleware
  m.Use(rest.DB())

  rest.Include(m)

  http.ListenAndServe(":8080", m)

}
