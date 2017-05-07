package main

import (
  "io"
  "fmt"
  "strconv"
  "encoding/csv"
  "strings"
)

func PlainOutput(data GithubDataPieces, writer io.Writer) error {
  for i, user := range data {
    fmt.Fprintf(writer, "#%+v: %+v (%+v):%+v (%+v)\n", i + 1, user.User.Name, user.User.Login, user.Contributions, strings.Join(user.Organizations, ","))
  }
  return nil
}

func CsvOutput(data GithubDataPieces, writer io.Writer) error {
  w := csv.NewWriter(writer)
  if err := w.Write([]string{"rank", "name", "login", "contributions", "organizations"}); err != nil {
    return err
  }
  for i, user := range data {
    rank := strconv.Itoa(i + 1)
    name := user.User.Name
    login := user.User.Login
    contribs := strconv.Itoa(user.Contributions)
    orgs := strings.Join(user.Organizations, ",")
    if err := w.Write([]string{ rank, name, login, contribs, orgs }); err != nil {
      return err
    }
  }
  w.Flush()
  return nil
}

type OutputFormat func(data GithubDataPieces, writer io.Writer) error
