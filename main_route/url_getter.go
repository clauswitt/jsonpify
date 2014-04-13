package main_route

import (
  "fmt"
  "bytes"
  "net/http"
)

func getUrl(url string) bytes.Buffer {
  response, err := http.Get( url )
  var err_buff bytes.Buffer
  if err != nil {
    err_buff.WriteString("{\"jsonp_error\": \"failed to open url\"}")
    return err_buff
  }
  buf := new(bytes.Buffer)
  buf.ReadFrom(response.Body)
  return *buf
}

func wrapContentFromUrl(url, wrapper string) string {
  content := getUrl(url)
  return fmt.Sprintf(wrapper, content.String())
}
