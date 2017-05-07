package main

import "flag"
import "fmt"
import "log"

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

  flag.Var(&locations, "location", "Location to query")
  flag.Parse()

  data, err := GithubTop(TopOptions { token: *token, locations: locations, amount: *amount })

  if err != nil {
    log.Fatal(err)
  }

  for i, user := range data {
    fmt.Printf("#%+v: %+v (%+v):%+v\n", i + 1, user.User.Name, user.User.Login, user.Contributions)
  }
}
