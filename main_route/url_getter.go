package main_route

import (
  "fmt"
  "bytes"
  "net/http"
  "github.com/bradfitz/gomemcache/memcache"
  "os"
  "crypto/md5"
  "io"
)

var mc *memcache.Client

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
  mc = memcache.New(os.Getenv("MEMCACHE_SERVERS"))
  h := md5.New()
  io.WriteString(h, url)
  md5_url := fmt.Sprintf("%x", h.Sum(nil))

  var content bytes.Buffer

  it, err := mc.Get(md5_url)


  if err == nil && it != nil {
    buf := new(bytes.Buffer)
    buf.Write(it.Value)
    content = *buf
  } else {
    content = getUrl(url)
    item := &memcache.Item{Key: md5_url, Value: []byte(content.Bytes())}
    mc.Set(item)
    buf := new(bytes.Buffer)
    buf.Write(item.Value)
    content = *buf
  }

  return fmt.Sprintf(wrapper, content.String())
}
