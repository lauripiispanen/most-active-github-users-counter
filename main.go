package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"github.com/lauripiispanen/most-active-github-users-counter/output"
	"github.com/lauripiispanen/most-active-github-users-counter/top"
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
	considerNum := flag.Int("consider", 1000, "Amount of users to consider")
	outputOpt := flag.String("output", "plain", "Output format: plain, csv")
	fileName := flag.String("file", "", "Output file (optional, defaults to stdout)")
	preset := flag.String("preset", "", "Preset (optional)")

	flag.Var(&locations, "location", "Location to query")
	flag.Parse()

	if *preset != "" {
		locations = Preset(*preset)
	}

	var format output.Format

	if *outputOpt == "plain" {
		format = output.PlainOutput
	} else if *outputOpt == "yaml" {
		format = output.YamlOutput
	} else if *outputOpt == "csv" {
		format = output.CsvOutput
	} else {
		log.Fatal("Unrecognized output format: ", *outputOpt)
	}

	data, err := top.GithubTop(top.Options{Token: *token, Locations: locations, Amount: *amount, ConsiderNum: *considerNum})

	if err != nil {
		log.Fatal(err)
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
