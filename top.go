package main

import "os"
import "fmt"
import "log"
import "sort"

const NUM_TOP = 512

func main() {
  var token string = os.Getenv("GITHUB_TOKEN")
  if token == "" {
    log.Fatal("Missing GITHUB token")
  }

  var client = NewGithubClient(DiskCache, TokenAuth(token))
/*  user, err := client.CurrentUser()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Printf("%+v\n", user)*/

  users, err := client.SearchUsers(UserSearchQuery{q: "location:finland+location:suomi+repos:>1+type:user", sort: "followers", order: "desc", per_page: 100, pages: 10})
  if err != nil {
    log.Fatal(err)
  }

  data := GithubDataPieces{}
  for _, username := range users {
    u, err := client.User(username)
    if err != nil {
      log.Fatal(err)
    }

    i, err := client.NumContributions(username)
    if err != nil {
      log.Fatal(err)
    }

    orgs, err := client.Organizations(username)
    if err != nil {
      log.Fatal(err)
    }

    data = append(data, GithubDataPiece{ User: u, Contributions: i, Organizations: orgs })
  }

  sort.Sort(data)

  for i, user := range data[:NUM_TOP] {
    fmt.Printf("#%+v: %+v (%+v):%+v\n", i + 1, user.User.Name, user.User.Login, user.Contributions)
  }
}

type GithubDataPiece struct {
  User          User
  Contributions int
  Organizations []string
}

type GithubDataPieces []GithubDataPiece

func (slice GithubDataPieces) Len() int {
    return len(slice)
}

func (slice GithubDataPieces) Less(i, j int) bool {
    return slice[i].Contributions > slice[j].Contributions
}

func (slice GithubDataPieces) Swap(i, j int) {
    slice[i], slice[j] = slice[j], slice[i]
}
