package main

import (
  "sort"
  "errors"
  "fmt"
  "regexp"
  "strings"
  "sync"
  "log"
  "time"
)

var companyLogin = regexp.MustCompile(`^\@([a-zA-Z0-9]+)$`)

func GithubTop(options TopOptions) (GithubDataPieces, error) {
  var token string = options.token
  if token == "" {
    return GithubDataPieces{}, errors.New("Missing GITHUB token")
  }

  var num_top = options.amount
  if num_top == 0 {
    num_top = 256
  }


  query := "repos:>1+type:user"
  for _, location := range locations {
    query = fmt.Sprintf("%s+location:%s", query, location)
  }

  var client = NewGithubClient(TokenAuth(token))
  users, err := client.SearchUsers(UserSearchQuery{q: query, sort: "followers", order: "desc", per_page: 100, pages: options.considerNum / 100})
  if err != nil {
    return GithubDataPieces{}, err
  }

  data := GithubDataPieces{}
  userContributions := make(UserContributions, 0)
  userContribChan := make(chan UserContribution)

  cachingClient := NewGithubClient(DiskCache, TokenAuth(token))


  var wg sync.WaitGroup
  wg.Add(len(users))

  throttle := time.Tick(time.Second / 10)

  for _, username := range users {
    go func(username string) {
      defer wg.Done()

      count, err := cachingClient.NumContributions(username)
      if err != nil {
        log.Fatal(err)
      }

      userContribChan <- UserContribution{ Username: username, Contributions: count }
    }(username)

    <- throttle
  }

  go func() {
      for userContrib := range userContribChan {
          userContributions = append(userContributions, userContrib)
      }
  }()

  wg.Wait()

  sort.Sort(userContributions)

  userContributions = userContributions[:num_top]


  pieces := make(chan GithubDataPiece)
  wg.Add(len(userContributions))

  throttle = time.Tick(time.Second / 10)

  for _, user := range userContributions {
    go func(user UserContribution) {
      defer wg.Done()
      u, err := cachingClient.User(user.Username)
      if err != nil {
        log.Fatal(err)
      }

      orgs, err := cachingClient.Organizations(user.Username)
      if err != nil {
        log.Fatal(err)
      }

      pieces <- GithubDataPiece{ User: u, Contributions: user.Contributions, Organizations: orgs }
    }(user)

    <- throttle
  }

  go func() {
      for piece := range pieces {
          data = append(data, piece)
      }
  }()

  wg.Wait()

  sort.Sort(data)

  data = data[:num_top]

  return data, nil
}

type UserContribution struct {
  Username      string
  Contributions int
}

type UserContributions []UserContribution

func (slice UserContributions) Len() int {
    return len(slice)
}

func (slice UserContributions) Less(i, j int) bool {
    return slice[i].Contributions > slice[j].Contributions
}

func (slice UserContributions) Swap(i, j int) {
    slice[i], slice[j] = slice[j], slice[i]
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
  considerNum int
}

type TopOrganization struct {
  Name        string
  MemberCount int
}

type TopOrganizations []TopOrganization

func (slice TopOrganizations) Len() int {
    return len(slice)
}

func (slice TopOrganizations) Less(i, j int) bool {
    return slice[i].MemberCount > slice[j].MemberCount
}

func (slice TopOrganizations) Swap(i, j int) {
    slice[i], slice[j] = slice[j], slice[i]
}


func (pieces GithubDataPieces) TopOrgs(count int) TopOrganizations {
  orgs_map := make(map[string]int)
  for _, piece := range pieces {
    user := piece.User
    user_orgs := piece.Organizations
    org_matches := companyLogin.FindStringSubmatch(strings.Trim(user.Company, " "))
    if (len(org_matches) > 0) {
      org_login := companyLogin.FindStringSubmatch(strings.Trim(user.Company, " "))[1]
      if (len(org_login) > 0 && !contains(user_orgs, org_login)) {
        user_orgs = append(user_orgs, org_login)
      }
    }

    for _, o := range user_orgs {
      orgs_map[o] = orgs_map[o] + 1
    }
  }

  orgs := TopOrganizations{}

  for k, v := range orgs_map {
    orgs = append(orgs, TopOrganization{ Name: k, MemberCount: v })
  }
  sort.Sort(orgs)
  if (len(orgs) > count) {
    return orgs[:count]
  } else {
    return orgs
  }

}

func contains (s []string, e string) bool {
  for _, a := range s {
    if a == e {
      return true
    }
  }
  return false
}
