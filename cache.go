package main

import "strings"
import "os"
import "fmt"
import "log"
import "io/ioutil"
import "net/http"

type producer func() ([]byte, error)

func CacheOnDisk(key string, f producer) ([]byte, error) {
  path := fmt.Sprintf(".cache/%s", key)
  if _, err := os.Stat(path); os.IsNotExist(err) {
    log.Printf("Value for key '%s' not found from cache", key)
    data, err := f()
    if err != nil {
      return []byte{}, err
    }
    err = ioutil.WriteFile(path, data, 0644)
    if err != nil {
      return []byte{}, err
    }
  }
  contents, err := ioutil.ReadFile(path)
  if err != nil {
    return []byte{}, err
  }
  return contents, nil
}

func DiskCache(r requester) requester {
  return func(req *http.Request) ([]byte, error) {
    path := req.URL.EscapedPath()
    key := strings.Replace(fmt.Sprintf("%s+%s", path[1:len(path)], req.URL.Query().Encode()), "/", "-", -1)
    return CacheOnDisk(key, func() ([]byte, error) {
      return r(req)
    })
  }
}
