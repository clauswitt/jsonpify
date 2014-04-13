package main_route

import (
  "net/http"
)


func UrlContentWrappedInCallback (res http.ResponseWriter, req *http.Request) string {

  url, callback := getParams(req)

  if callback  == nil {
    return "{\"jsonp_error\": \"missing callback parameter\"}"
  }

  if url != nil {
    return wrapContent(url[0], callback[0])
  } else {
    // could be changed to show landing page
    return "{\"jsonp_error\": \"missing url parameter\"}"
  }
}

