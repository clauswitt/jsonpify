package main_route

import (
  "bytes"
  "net/http"
)

func GetUrl(url string) bytes.Buffer {
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

func GetParams(req *http.Request) ([]string, []string) {
  params :=  req.URL.Query()
  return params["url"], params["callback"]
}

func WrapContent(url, callback string) string {

  var buffer bytes.Buffer
  var content bytes.Buffer
  content = GetUrl(url)
  buffer.WriteString(callback)
  buffer.WriteString("(")
  buffer.Write(content.Bytes())
  buffer.WriteString(")")
  return buffer.String()

}
