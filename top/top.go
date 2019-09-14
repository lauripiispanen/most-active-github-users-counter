package top

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/lauripiispanen/most-active-github-users-counter/github"
	"github.com/lauripiispanen/most-active-github-users-counter/net"
)

var companyLogin = regexp.MustCompile(`^\@([a-zA-Z0-9]+)$`)

func GithubTop(options Options) (GithubDataPieces, error) {
	var token = options.Token
	if token == "" {
		return GithubDataPieces{}, errors.New("Missing GITHUB token")
	}

	var numTop = options.Amount
	if numTop == 0 {
		numTop = 256
	}

	query := "repos:>1 type:user"
	for _, location := range options.Locations {
		query = fmt.Sprintf("%s location:%s", query, location)
	}

	var client = github.NewGithubClient(net.TokenAuth(token))
	users, err := client.SearchUsers(github.UserSearchQuery{Q: query, Sort: "followers", Order: "desc", MaxUsers: options.ConsiderNum})
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
			pieces <- GithubDataPiece{User: user, Contributions: count}
		}(user)

		<-throttle
	}

	for piece := range pieces {
		data = append(data, piece)
		if len(data) >= len(users) {
			close(pieces)
		}
	}

	sort.Sort(data)
	if len(data) < numTop {
		numTop = len(data)
	}
	data = data[:numTop]

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

type Options struct {
	Token       string
	Locations   []string
	Amount      int
	ConsiderNum int
}

type Organization struct {
	Name        string
	MemberCount int
}

type Organizations []Organization

func (slice Organizations) Len() int {
	return len(slice)
}

func (slice Organizations) Less(i, j int) bool {
	return slice[i].MemberCount > slice[j].MemberCount
}

func (slice Organizations) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (slice GithubDataPieces) TopOrgs(count int) Organizations {
	orgsMap := make(map[string]int)
	for _, piece := range slice {
		user := piece.User
		userOrgs := user.Organizations
		orgMatches := companyLogin.FindStringSubmatch(strings.Trim(user.Company, " "))
		if len(orgMatches) > 0 {
			orgLogin := companyLogin.FindStringSubmatch(strings.Trim(user.Company, " "))[1]
			if len(orgLogin) > 0 && !contains(userOrgs, orgLogin) {
				userOrgs = append(userOrgs, orgLogin)
			}
		}

		for _, o := range userOrgs {
			org := strings.ToLower(o)
			orgsMap[org] = orgsMap[org] + 1
		}
	}

	orgs := Organizations{}

	for k, v := range orgsMap {
		orgs = append(orgs, Organization{Name: k, MemberCount: v})
	}
	sort.Sort(orgs)
	if len(orgs) > count {
		return orgs[:count]
	}
	return orgs
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
