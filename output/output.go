package output

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/lauripiispanen/most-active-github-users-counter/top"
)

type Format func(users top.GithubUsers, writer io.Writer) error

func PlainOutput(users top.GithubUsers, writer io.Writer) error {
	fmt.Fprintln(writer, "USERS\n--------")
	for i, user := range users {
		fmt.Fprintf(writer, "#%+v: %+v (%+v):%+v (%+v) %+v\n", i+1, user.Name, user.Login, user.ContributionCount, user.Company, strings.Join(user.Organizations, ","))
	}
	fmt.Fprintln(writer, "\nORGANIZATIONS\n--------")
	for i, org := range users.TopOrgs(10) {
		fmt.Fprintf(writer, "#%+v: %+v (%+v)\n", i+1, org.Name, org.MemberCount)
	}
	return nil
}

func CsvOutput(users top.GithubUsers, writer io.Writer) error {
	w := csv.NewWriter(writer)
	if err := w.Write([]string{"rank", "name", "login", "contributions", "company", "organizations"}); err != nil {
		return err
	}
	for i, user := range users {
		rank := strconv.Itoa(i + 1)
		name := user.Name
		login := user.Login
		contribs := strconv.Itoa(user.ContributionCount)
		orgs := strings.Join(user.Organizations, ",")
		company := user.Company
		if err := w.Write([]string{rank, name, login, contribs, company, orgs}); err != nil {
			return err
		}
	}
	w.Flush()
	return nil
}

func YamlOutput(users top.GithubUsers, writer io.Writer) error {
	fmt.Fprintln(writer, "users:")
	for i, user := range users {
		fmt.Fprintf(
			writer,
			`
  - rank: %+v
    name: '%+v'
    login: '%+v'
    avatarUrl: '%+v'
    contributions: %+v
    company: '%+v'
    organizations: '%+v'
`,
			i+1,
			strings.Replace(user.Name, "'", "''", -1),
			strings.Replace(user.Login, "'", "''", -1),
			user.AvatarURL,
			user.ContributionCount,
			strings.Replace(user.Company, "'", "''", -1),
			strings.Replace(strings.Join(user.Organizations, ","), "'", "''", -1))
	}
	fmt.Fprintln(writer, "\norganizations:")

	for i, org := range users.TopOrgs(10) {
		fmt.Fprintf(
			writer,
			`
  - rank: %+v
    name: '%+v'
    membercount: %+v
`,
			i+1,
			strings.Replace(org.Name, "'", "''", -1),
			org.MemberCount)
	}

	fmt.Fprintf(writer, "generated: %+v\n", time.Now())

	return nil
}
