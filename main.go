package main

import (
  "flag"
  "log"
  "bufio"
  "os"
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
  } else if *output == "yaml" {
    format = YamlOutput
  } else if *output == "csv" {
    format = CsvOutput
  }

  var writer *bufio.Writer
  if *fileName != "" {
    f, err := os.Create(*fileName)
    if err != nil {
      log.Fatal(err)
    }
    writer = bufio.NewWriter(f)
    defer f.Close()
  } else {
     writer = bufio.NewWriter(os.Stdout)
  }

  err = format(data, writer)
  if err != nil {
    log.Fatal(err)
  }
  writer.Flush()
}
