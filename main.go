package main

import (
  "flag"
  "fmt"
  "strconv"
  "log"
  "io"
  "bufio"
  "os"
  "encoding/csv"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
    return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
    *i = append(*i, value)
    return nil
}

var locations arrayFlags

func main() {
  token := flag.String("token", "", "Github auth token")
  amount := flag.Int("amount", 256, "Amount of users to show")
  output := flag.String("output", "plain", "Output format: plain, csv")
  fileName := flag.String("file", "", "Output file (optional, defaults to stdout)")

  flag.Var(&locations, "location", "Location to query")
  flag.Parse()

  data, err := GithubTop(TopOptions { token: *token, locations: locations, amount: *amount })

  if err != nil {
    log.Fatal(err)
  }

  var format OutputFormat

  if *output == "plain" {
    format = PlainOutput
  } else if *output == "csv" {
    format = CsvOutput
  }

  var writer io.Writer
  if *fileName != "" {
    f, err := os.Create(*fileName)
    if err != nil {
      log.Fatal(err)
    }
    defer f.Close()
    writer = bufio.NewWriter(f)
  } else {
     writer = os.Stdout
  }

  format(data, writer)
}

func PlainOutput(data GithubDataPieces, writer io.Writer) {
  for i, user := range data {
    fmt.Fprintf(writer, "#%+v: %+v (%+v):%+v\n", i + 1, user.User.Name, user.User.Login, user.Contributions)
  }
}

func CsvOutput(data GithubDataPieces, writer io.Writer) {
  w := csv.NewWriter(writer)
  if err := w.Write([]string{"name", "login", "contributions"}); err != nil {
    log.Fatal(err)
  }
  for _, user := range data {
    if err := w.Write([]string{ user.User.Name, user.User.Login, strconv.Itoa(user.Contributions) }); err != nil {
      log.Fatal(err)
    }
  }
  w.Flush()
}

type OutputFormat func(data GithubDataPieces, writer io.Writer)
