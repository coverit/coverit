package main

import (
  "fmt"

  "github.com/codegangsta/martini"
  "github.com/codegangsta/martini-contrib/render"
  "github.com/tmilewski/goenv"

  "./rest"
)

func init() {
    err := goenv.Load()

    if err != nil {
        fmt.Println(err)
    }
}


func main() {
  m := martini.Classic()
  m.Use(render.Renderer(render.Options {
    // nothing here to configure
  }))
  // use the Mongo middleware
  m.Use(rest.DB())

  rest.Include(m)

  m.Run()
}
