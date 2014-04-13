package main

import (
  "github.com/clauswitt/jsonpify/main_route"
  "github.com/go-martini/martini"
)

func main() {
  m := martini.Classic()
  m.Get("/", main_route.UrlContentWrappedInCallback)
  m.Run()
}
