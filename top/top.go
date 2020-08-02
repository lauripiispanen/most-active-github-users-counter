package top

import (
	"errors"
	"fmt"

	"most-active-github-users-counter/github"
	"most-active-github-users-counter/net"
)

func GithubTop(options Options) (github.GithubSearchResults, error) {
	var token = options.Token
	if token == "" {
		return github.GithubSearchResults{}, errors.New("Missing GITHUB token")
	}

	query := "type:user"
	for _, location := range options.Locations {
		query = fmt.Sprintf("%s location:%s", query, location)
	}

	for _, location := range options.ExcludeLocations {
		query = fmt.Sprintf("%s -location:%s", query, location)
	}

	var client = github.NewGithubClient(net.TokenAuth(token))
	users, err := client.SearchUsers(github.UserSearchQuery{Q: query, Sort: "followers", Order: "desc", MaxUsers: options.ConsiderNum})
	if err != nil {
		return github.GithubSearchResults{}, err
	}
	return users, nil
}

type Options struct {
	Token            string
	Locations        []string
	ExcludeLocations []string
	Amount           int
	ConsiderNum      int
}
