package output

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"most-active-github-users-counter/github"
	"most-active-github-users-counter/top"
)

type Format func(results github.GithubSearchResults, writer io.Writer, options top.Options) error

func PlainOutput(results github.GithubSearchResults, writer io.Writer, options top.Options) error {
	users := GithubUserList(results.Users)
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

func CsvOutput(results github.GithubSearchResults, writer io.Writer, options top.Options) error {
	users := GithubUserList(results.Users)
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

type ContributionsSelector func(github.User) int

func selectCommits(user github.User) int {
	return user.CommitsCount
}

func selectContributions(user github.User) int {
	return user.ContributionCount
}

func selectPublicContributions(user github.User) int {
	return user.PublicContributionCount
}

func YamlOutput(results github.GithubSearchResults, writer io.Writer, options top.Options) error {
	users := GithubUserList(results.Users)
	outputUsers := func(user []github.User, cs ContributionsSelector) {
		for i, u := range user {
			contributionCount := cs(u)
			fmt.Fprintf(
				writer,
				`
  - rank: %+v
    name: %+v
    login: %+v
    avatarUrl: %+v
    contributions: %+v
    company: %+v
    organizations: %+v
`,
				i+1,
				strconv.QuoteToASCII(u.Name),
				strconv.QuoteToASCII(u.Login),
				u.AvatarURL,
				contributionCount,
				strconv.QuoteToASCII(u.Company),
				strconv.QuoteToASCII(strings.Join(u.Organizations, ",")))
		}
	}

	topPublic := users.TopCommits(options.Amount)
	fmt.Fprintln(writer, "users:")
	outputUsers(topPublic, selectCommits)

	topContributions := users.TopPublic(options.Amount)
	fmt.Fprintln(writer, "users_public_contributions:")
	outputUsers(topContributions, selectPublicContributions)

	topPrivate := users.TopPrivate(options.Amount)
	fmt.Fprintln(writer, "\nprivate_users:")
	outputUsers(topPrivate, selectContributions)

	outputOrganizations := func(orgs Organizations) {
		for i, org := range orgs {
			fmt.Fprintf(
				writer,
				`
  - rank: %+v
    name: %+v
    membercount: %+v
`,
				i+1,
				strconv.QuoteToASCII(org.Name),
				org.MemberCount)
		}
	}

	fmt.Fprintln(writer, "\norganizations:")
	outputOrganizations(topPublic.TopOrgs(10))
	fmt.Fprintln(writer, "\npublic_contributions_organizations:")
	outputOrganizations(topContributions.TopOrgs(10))
	fmt.Fprintln(writer, "\nprivate_organizations:")
	outputOrganizations(topPrivate.TopOrgs(10))

	fmt.Fprintf(writer, "generated: %+v\n", time.Now())
	fmt.Fprintf(writer, "min_followers_required: %+v\n", results.MinimumFollowerCount)
	fmt.Fprintf(writer, "total_user_count: %+v\n", results.TotalUserCount)

	return nil
}

var companyLogin = regexp.MustCompile(`^\@([a-zA-Z0-9]+)$`)

func trim(users GithubUserList, numTop int) GithubUserList {
	if numTop == 0 {
		numTop = 256
	}
	if len(users) < numTop {
		numTop = len(users)
	}
	return users[:numTop]
}

func clone(users GithubUserList) GithubUserList {
	usersCloned := make(GithubUserList, len(users))
	copy(usersCloned, users)
	return usersCloned
}

type GithubUserList []github.User

func (users GithubUserList) TopCommits(amount int) GithubUserList {
	u := TopCommitsUsers(clone(users))
	sort.Sort(u)
	return trim(GithubUserList(u), amount)
}

func (users GithubUserList) TopPublic(amount int) GithubUserList {
	u := TopPublicUsers(clone(users))
	sort.Sort(u)
	return trim(GithubUserList(u), amount)
}

func (users GithubUserList) TopPrivate(amount int) GithubUserList {
	u := TopPrivateUsers(clone(users))
	sort.Sort(u)
	return trim(GithubUserList(u), amount)
}

func (slice GithubUserList) MinFollowers() int {
	if len(slice) == 0 {
		return 0
	}
	followers := math.MaxInt32
	for _, user := range slice {
		if user.FollowerCount < followers {
			followers = user.FollowerCount
		}
	}
	return followers
}

type TopCommitsUsers GithubUserList

func (slice TopCommitsUsers) Len() int {
	return len(slice)
}

func (slice TopCommitsUsers) Less(i, j int) bool {
	return slice[i].CommitsCount > slice[j].CommitsCount
}

func (slice TopCommitsUsers) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

type TopPublicUsers GithubUserList

func (slice TopPublicUsers) Len() int {
	return len(slice)
}

func (slice TopPublicUsers) Less(i, j int) bool {
	return slice[i].PublicContributionCount > slice[j].PublicContributionCount
}

func (slice TopPublicUsers) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

type TopPrivateUsers GithubUserList

func (slice TopPrivateUsers) Len() int {
	return len(slice)
}

func (slice TopPrivateUsers) Less(i, j int) bool {
	return slice[i].ContributionCount > slice[j].ContributionCount
}

func (slice TopPrivateUsers) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
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

func (slice GithubUserList) TopOrgs(count int) Organizations {
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
