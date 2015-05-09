package rest

import (
  "os"
  "github.com/codegangsta/martini"
  "labix.org/v2/mgo"
)

/*
   the function returns a martini.Handler which is called on each request. We simply clone
   the session for each request and close it when the request is complete. The call to c.Map
   maps an instance of *mgo.Database to the request context. Then *mgo.Database
   is injected into each handler function.
*/
func DB() martini.Handler {
  session, err := mgo.Dial(os.Getenv("MONGO_URL")) // mongodb://localhost
  if err != nil {
    panic(err)
  }

  return func(c martini.Context) {
    s := session.Clone()
    c.Map(s.DB(os.Getenv("MONGO_DB"))) // coverit
    defer s.Close()
    c.Next()
  }
}
