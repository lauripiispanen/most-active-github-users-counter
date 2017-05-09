package main

import (
  "io"
  "fmt"
  "strconv"
  "encoding/csv"
  "strings"
)

type OutputFormat func(data GithubDataPieces, writer io.Writer) error

func PlainOutput(data GithubDataPieces, writer io.Writer) error {
  for i, piece := range data {
    fmt.Fprintf(writer, "#%+v: %+v (%+v):%+v (%+v) %+v\n", i + 1, piece.User.Name, piece.User.Login, piece.Contributions, piece.User.Company, strings.Join(piece.Organizations, ","))
  }
  return nil
}

func CsvOutput(data GithubDataPieces, writer io.Writer) error {
  w := csv.NewWriter(writer)
  if err := w.Write([]string{"rank", "name", "login", "contributions", "company", "organizations"}); err != nil {
    return err
  }
  for i, piece := range data {
    rank := strconv.Itoa(i + 1)
    name := piece.User.Name
    login := piece.User.Login
    contribs := strconv.Itoa(piece.Contributions)
    orgs := strings.Join(piece.Organizations, ",")
    company := piece.User.Company
    if err := w.Write([]string{ rank, name, login, contribs, company, orgs }); err != nil {
      return err
    }
  }
  w.Flush()
  return nil
}

func YamlOutput(data GithubDataPieces, writer io.Writer) error {
  fmt.Fprintln(writer, "users:")
  for i, piece := range data {
    fmt.Fprintf(
      writer,
      `
  - rank: %+v
    name: %+v
    login: %+v
    id: %+v
    contributions: %+v
    company: %+v
    organizations: %+v
      `,
      i + 1,
      piece.User.Name,
      piece.User.Login,
      piece.User.Id,
      piece.Contributions,
      piece.User.Company,
      strings.Join(piece.Organizations, ","))
  }
  return nil
}
