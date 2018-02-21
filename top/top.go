package top

import (
  "sort"
  "errors"
  "fmt"
  "regexp"
  "strings"
  "log"
  "time"
  "github.com/lauripiispanen/most-active-github-users-counter/github"
  "github.com/lauripiispanen/most-active-github-users-counter/net"
)

var companyLogin = regexp.MustCompile(`^\@([a-zA-Z0-9]+)$`)

func GithubTop(options TopOptions) (GithubDataPieces, error) {
  var token string = options.Token
  if token == "" {
    return GithubDataPieces{}, errors.New("Missing GITHUB token")
  }

  var num_top = options.Amount
  if num_top == 0 {
    num_top = 256
  }


  query := "repos:>1 type:user"
  for _, location := range options.Locations {
    query = fmt.Sprintf("%s location:%s", query, location)
  }

  var client = github.NewGithubClient(net.TokenAuth(token))
  users, err := client.SearchUsers(github.UserSearchQuery{Q: query, Sort: "followers", Order: "desc", Per_page: 100, Pages: options.ConsiderNum / 100})
  if err != nil {
    return GithubDataPieces{}, err
  }


  data := GithubDataPieces{}

  pieces := make(chan GithubDataPiece)
  throttle := time.Tick(time.Second / 20)

  for _, user := range users {
    go func(user github.User) {
      count, err := client.NumContributions(user.Login)
      if err != nil {
        log.Fatal(err)
      }
      pieces <- GithubDataPiece{ User: user, Contributions: count }
    }(user)

    <- throttle
  }

  for piece := range pieces {
    data = append(data, piece)
    if (len(data) >= len(users)) {
      close(pieces)
    }
  }

  sort.Sort(data)
  if (len(data) < num_top) {
    num_top = len(data)
  }
  data = data[:num_top]

  return data, nil
}

type GithubDataPiece struct {
  User          github.User
  Contributions int
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
  Token     string
  Locations []string
  Amount    int
  ConsiderNum int
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
    user_orgs := user.Organizations
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
