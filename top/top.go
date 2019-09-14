package top

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/lauripiispanen/most-active-github-users-counter/github"
	"github.com/lauripiispanen/most-active-github-users-counter/net"
)

var companyLogin = regexp.MustCompile(`^\@([a-zA-Z0-9]+)$`)

func GithubTop(options Options) (GithubUsers, error) {
	var token = options.Token
	if token == "" {
		return GithubUsers{}, errors.New("Missing GITHUB token")
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
		return GithubUsers{}, err
	}

	sort.Sort(GithubUsers(users))
	if len(users) < numTop {
		numTop = len(users)
	}
	users = users[:numTop]

	return users, nil
}

type GithubUsers []github.User

func (slice GithubUsers) Len() int {
	return len(slice)
}

func (slice GithubUsers) Less(i, j int) bool {
	return slice[i].ContributionCount > slice[j].ContributionCount
}

func (slice GithubUsers) Swap(i, j int) {
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

func (slice GithubUsers) TopOrgs(count int) Organizations {
	orgsMap := make(map[string]int)
	for _, user := range slice {
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
