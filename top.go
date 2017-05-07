package main

import (
  "sort"
  "errors"
  "fmt"
  "sync"
  "log"
)

func GithubTop(options TopOptions) (GithubDataPieces, error) {
  var token string = options.token
  if token == "" {
    return GithubDataPieces{}, errors.New("Missing GITHUB token")
  }

  var num_top = options.amount
  if num_top == 0 {
    num_top = 256
  }

  var client = NewGithubClient(DiskCache, TokenAuth(token))

  query := "repos:>1+type:user"
  for _, location := range locations {
    query = fmt.Sprintf("%s+location:%s", query, location)
  }


  users, err := client.SearchUsers(UserSearchQuery{q: query, sort: "followers", order: "desc", per_page: 100, pages: 10})
  if err != nil {
    return GithubDataPieces{}, err
  }

  data := GithubDataPieces{}
  pieces := make(chan GithubDataPiece)

  var wg sync.WaitGroup
  wg.Add(len(users))

  for _, username := range users {
    go func(username string) {
      defer wg.Done()
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

      pieces <- GithubDataPiece{ User: u, Contributions: i, Organizations: orgs }
    }(username)
  }

  go func() {
      for piece := range pieces {
          data = append(data, piece)
      }
  }()

  wg.Wait()

  sort.Sort(data)

  return data[:num_top], nil
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

type TopOptions struct {
  token     string
  locations []string
  amount    int
}
