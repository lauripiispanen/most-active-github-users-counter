package main

import "net/http"
import "fmt"
import "io/ioutil"

type requester func(req *http.Request) ([]byte, error)

type wrapper func(requester) requester

func compose(wrappers ...wrapper) wrapper {
  return func(r requester) requester {
    for _, wrapper := range wrappers {
      r = wrapper(r)
    }
    return r
  }
}

func TokenAuth(token string) wrapper {
  return func(r requester) requester {
    return func(req *http.Request) ([]byte, error) {
      req.Header.Add("Authorization", fmt.Sprintf("token %s", token))
      return r(req)
    }
  }
}

func Requester(client *http.Client) requester {
  return func(req *http.Request) ([]byte, error) {
    resp, err := client.Do(req)

    if err != nil {
      return []byte{}, err
    }

    bodyText, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      return []byte{}, err
    }

    return bodyText, nil
  }
}
