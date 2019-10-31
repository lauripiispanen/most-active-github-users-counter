package github

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/lauripiispanen/most-active-github-users-counter/net"
)

const root string = "https://api.github.com/"

type HTTPGithubClient struct {
	wrappers []net.Wrapper
}

func (client HTTPGithubClient) Request(url string, body string) ([]byte, error) {
	httpClient := &http.Client{}
	var req *http.Request
	var err error
	if body != "" {
		req, err = http.NewRequest("POST", url, strings.NewReader(body))
	} else {
		req, err = http.NewRequest("GET", url, nil)
	}

	if err != nil {
		return []byte{}, err
	}

	return net.Compose(client.wrappers...)(net.MakeRequester(httpClient))(req)
}

func (client HTTPGithubClient) CurrentUser() (User, error) {
	body, err := client.Request(fmt.Sprintf("%suser", root), "")
	if err != nil {
		return User{}, err
	}

	user := User{}
	if err := json.Unmarshal(body, &user); err != nil {
		return User{}, err
	}
	return user, nil
}

func (client HTTPGithubClient) User(login string) (User, error) {
	body, err := client.Request(fmt.Sprintf("%susers/%s", root, login), "")
	if err != nil {
		return User{}, err
	}

	user := User{}
	if err := json.Unmarshal(body, &user); err != nil {
		return User{}, err
	}
	return user, nil
}

func (client HTTPGithubClient) SearchUsers(query UserSearchQuery) ([]User, error) {
	users := []User{}
	userLogins := map[string]bool{}

	totalCount := 0
	minFollowerCount := -1
	maxPerQuery := 1000
	perPage := 5

Pages:
	for totalCount < query.MaxUsers {
		previousCursor := ""
		followerCountQueryStr := ""
		if minFollowerCount >= 0 {
			followerCountQueryStr = fmt.Sprintf(" followers:<%d", minFollowerCount)
		}
		for currentPage := 1; currentPage <= (maxPerQuery / perPage); currentPage++ {
			cursorQueryStr := ""
			if previousCursor != "" {
				cursorQueryStr = fmt.Sprintf(", after: \\\"%s\\\"", previousCursor)
			}
			graphQlString := fmt.Sprintf(`{ "query": "query {
        search(type: USER, query:\"%s%s sort:%s-%s\", first: %d%s) {
          edges {
            node {
              __typename
              ... on User {
                login,
                avatarUrl,
                name,
                company,
                organizations(first: 100) {
                  nodes {
                    login
                  }
                }
                followers {
                  totalCount
                }
                contributionsCollection {
                  contributionCalendar {
                    totalContributions
                  }
                  restrictedContributionsCount
                }
              }
            },
            cursor
          }
        }
      }" }`, query.Q, followerCountQueryStr, query.Sort, query.Order, perPage, cursorQueryStr)

			re := regexp.MustCompile(`\r?\n`)
			graphQlString = re.ReplaceAllString(graphQlString, " ")

			body, err := client.Request("https://api.github.com/graphql", graphQlString)
			if err != nil {
				log.Fatal(err)
			}

			var response interface{}
			if err := json.Unmarshal(body, &response); err != nil {
				log.Fatal(err)
			}
			rootNode := response.(map[string]interface{})
			if val, ok := rootNode["errors"]; ok {
				log.Fatalf("%s", val)
			}
			dataNode := rootNode["data"].(map[string]interface{})
			searchNode := dataNode["search"].(map[string]interface{})
			edgeNodes := searchNode["edges"].([]interface{})

			if len(edgeNodes) == 0 {
				break Pages
			}
			totalCount += len(edgeNodes)

		Edges:
			for _, edge := range edgeNodes {
				edgeNode := edge.(map[string]interface{})
				userNode := edgeNode["node"].(map[string]interface{})
				typename := userNode["__typename"].(string)
				if typename != "User" {
					continue Edges
				}
				login := userNode["login"].(string)
				avatarURL := userNode["avatarUrl"].(string)
				name := strPropOrEmpty(userNode, "name")
				company := strPropOrEmpty(userNode, "company")
				organizations := []string{}

				orgNodes := userNode["organizations"].(map[string]interface{})["nodes"].([]interface{})
				for _, orgNode := range orgNodes {

					organizations = append(organizations, orgNode.(map[string]interface{})["login"].(string))
				}

				followerCount := int(userNode["followers"].(map[string]interface{})["totalCount"].(float64))
				contributionsCollection := userNode["contributionsCollection"].(map[string]interface{})
				contributionCount := int(contributionsCollection["contributionCalendar"].(map[string]interface{})["totalContributions"].(float64))
				privateContributionCount := int(contributionsCollection["restrictedContributionsCount"].(float64))

				user := User{
					Login:                    login,
					AvatarURL:                avatarURL,
					Name:                     name,
					Company:                  company,
					Organizations:            organizations,
					FollowerCount:            followerCount,
					ContributionCount:        contributionCount,
					PublicContributionCount:  (contributionCount - privateContributionCount),
					PrivateContributionCount: privateContributionCount}

				if !userLogins[login] {
					userLogins[login] = true
					users = append(users, user)
				}

				previousCursor = edgeNode["cursor"].(string)
				minFollowerCount = int(followerCount)
			}
		}
	}

	return users, nil
}

func strPropOrEmpty(obj map[string]interface{}, prop string) string {
	switch t := obj[prop].(type) {
	case string:
		return t
	default:
		return ""
	}

}

func (client HTTPGithubClient) Organizations(login string) ([]string, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/orgs", login)
	body, err := client.Request(url, "")
	if err != nil {
		log.Fatalf("error requesting organizations for user %+v", login)
		return []string{}, err
	}
	orgResp := []OrgResponse{}
	err = json.Unmarshal(body, &orgResp)
	if err != nil {
		log.Fatalf("error parsing organizations JSON for user %+v", login)
		return []string{}, err
	}
	orgs := []string{}

	for _, org := range orgResp {
		orgs = append(orgs, org.Organization)
	}

	return orgs, err
}

type OrgResponse struct {
	Organization string `json:"login"`
}

func NewGithubClient(wrappers ...net.Wrapper) HTTPGithubClient {
	return HTTPGithubClient{wrappers: wrappers}
}

type User struct {
	Login                    string
	AvatarURL                string
	Name                     string
	Company                  string
	Organizations            []string
	FollowerCount            int
	ContributionCount        int
	PublicContributionCount  int
	PrivateContributionCount int
}

type UserSearchQuery struct {
	Q        string
	Sort     string
	Order    string
	MaxUsers int
}
