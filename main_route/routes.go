package main_route

import (
  "fmt"
  "net/http"
)

func getParams(req *http.Request) ([]string, []string) {
  params :=  req.URL.Query()
  return params["url"], params["callback"]
}

func UrlContentWrappedInCallback (res http.ResponseWriter, req *http.Request) string {

  url, callback := getParams(req)

  if callback  == nil {
    return "{\"jsonp_error\": \"missing callback parameter\"}"
  }

  if url != nil {
    wrapString := fmt.Sprintf("%v(%%v)", callback[0])
    return wrapContentFromUrl(url[0], wrapString)
  } else {
    // could be changed to show landing page
    return "{\"jsonp_error\": \"missing url parameter\"}"
  }
}

