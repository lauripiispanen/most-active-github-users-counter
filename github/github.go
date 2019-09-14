package github

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/anaskhan96/soup"
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
	pages := 1
	if query.Pages > 0 {
		pages = query.Pages
	}
	if pages > 10 {
		pages = 10
	}

	users := []User{}

	currentPage := 1
	totalCount := 0
	previousCursor := ""
	cursorQueryStr := ""
	for currentPage <= pages {
		if previousCursor != "" {
			cursorQueryStr = fmt.Sprintf(", after: \\\"%s\\\"", previousCursor)
		}
		graphQlString := fmt.Sprintf(`{ "query": "query {
      search(type: USER, query:\"%s sort:%s-%s\", first: %d%s) {
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
            }
          },
          cursor
        }
      }
    }" }`, query.Q, query.Sort, query.Order, query.PerPage, cursorQueryStr)

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
		totalCount += len(edgeNodes)
		for _, edge := range edgeNodes {
			edgeNode := edge.(map[string]interface{})
			userNode := edgeNode["node"].(map[string]interface{})
			login := userNode["login"].(string)
			avatarURL := userNode["avatarUrl"].(string)
			name := strPropOrEmpty(userNode, "name")
			company := strPropOrEmpty(userNode, "company")
			organizations := []string{}

			orgNodes := userNode["organizations"].(map[string]interface{})["nodes"].([]interface{})
			for _, orgNode := range orgNodes {

				organizations = append(organizations, orgNode.(map[string]interface{})["login"].(string))
			}

			user := User{Login: login, AvatarURL: avatarURL, Name: name, Company: company, Organizations: organizations}
			users = append(users, user)

			previousCursor = edgeNode["cursor"].(string)
		}
		currentPage++
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

func (client HTTPGithubClient) NumContributions(login string) (int, error) {
	body, err := client.Request(fmt.Sprintf("https://github.com/users/%s/contributions", login), "")
	if err != nil {
		log.Fatalf("error requesting contributions for user %+v: %+v", login, err)
		return -1, err
	}
	doc := soup.HTMLParse(string(body))
	dayNodes := doc.FindAll("rect", "class", "day")

	count := 0
	for _, dayNode := range dayNodes {
		i, err := strconv.Atoi(dayNode.Attrs()["data-count"])
		if err != nil {
			return -1, err
		}
		count += int(i)
	}

	return count, err
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
	Login         string
	AvatarURL     string
	Name          string
	Company       string
	Organizations []string
}

type UserSearchQuery struct {
	Q       string
	Sort    string
	Order   string
	PerPage int
	Pages   int
}
