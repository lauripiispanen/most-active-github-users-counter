package main

import (
  "strconv"
  "fmt"
  "net/http"
  "net/url"
  "encoding/json"
  "encoding/xml"
  "log"
  "time"
)

const root string = "https://api.github.com/"

type HttpGithubClient struct {
  wrappers []wrapper
}

func (c HttpGithubClient) Request(url string) ([]byte, error) {
  client := &http.Client {}
  req, err := http.NewRequest("GET", url, nil)

  if err != nil {
    return []byte{}, err
  }

  return compose(c.wrappers...)(Requester(client))(req)
}

func (client HttpGithubClient) CurrentUser() (User, error) {
  body, err := client.Request(fmt.Sprintf("%suser", root))
  if err != nil {
    return User {}, err
  }

  user := User {}
  if err := json.Unmarshal(body, &user); err != nil {
    return User {}, err
  }
  return user, nil
}

func (client HttpGithubClient) User(login string) (User, error) {
  body, err := client.Request(fmt.Sprintf("%susers/%s", root, login))
  if err != nil {
    return User {}, err
  }

  user := User {}
  if err := json.Unmarshal(body, &user); err != nil {
    return User {}, err
  }
  return user, nil
}

func (client HttpGithubClient) SearchUsers(query UserSearchQuery) ([]string, error) {
  v := url.Values {}
  v.Set("q", query.q)
  v.Set("sort", query.sort)
  v.Set("order", query.order)
  if query.per_page > 0 {
    v.Set("per_page", strconv.Itoa(query.per_page))
  }
  pages := 1
  if query.pages > 0 {
    pages = query.pages
  }
  if pages > 10 {
    pages = 10
  }

  logins := []string{}
  currentPage := 1
  total_count := 0
  max_tries_per_page := 10

  throttle := time.Tick(time.Second * 3)

  for currentPage <= pages {
    items := make([]interface{}, 0)

    CURRENT_PAGE_ATTEMPT:
    for currentTry := 0; currentTry < max_tries_per_page; currentTry++ {
      localValues := url.Values {}
      for k,v := range v {
        localValues[k] = v
      }
      localValues.Set("page", strconv.Itoa(currentPage))
      q, err := url.QueryUnescape(localValues.Encode())
      if err != nil {
        log.Fatal(err)
      }

      url := fmt.Sprintf("%ssearch/users?%s", root, q)

      fmt.Printf("%s\n", url)

      body, err := client.Request(url)
      if err != nil {
        log.Fatal(err)
      }

      var response interface {}
      if err := json.Unmarshal(body, &response); err != nil {
        log.Fatal(err)
      }
      m := response.(map[string]interface{})
      if m["total_count"] == nil {
        fmt.Printf("Total count was nil for page %+v", currentPage)
        continue CURRENT_PAGE_ATTEMPT
      }

      total := int(m["total_count"].(float64))
      if (total >= total_count) {
        total_count = total
        items = m["items"].([]interface{})

        fmt.Printf("Established total count %+v for page %+v\n", total_count, currentPage)
        if (currentPage > 1) {
          break CURRENT_PAGE_ATTEMPT
        }
      }
      <- throttle
    }

    for _, item := range items {
      login := item.(map[string]interface{})["login"].(string)
      logins = append(logins, login)
    }

    currentPage += 1
  }

  return logins, nil
}

type ContributionsSvgRoot struct {
  G struct {
    G []struct {
      Rect []struct {
        Count string `xml:"data-count,attr"`
      } `xml:"rect"`
    } `xml:"g"`
  } `xml:"g"`
}

func (client HttpGithubClient) NumContributions(login string) (int, error) {
  body, err := client.Request(fmt.Sprintf("https://github.com/users/%s/contributions", login))
  if err != nil {
    log.Fatalf("error requesting contributions for user %+v", login)
    return -1, err
  }
  graph := ContributionsSvgRoot {}
  err = xml.Unmarshal(body, &graph)
  count := 0
  for _, s := range graph.G.G {
    for _, r := range s.Rect {
      i, err := strconv.Atoi(r.Count)
      if err != nil {
        return -1, err
      }
      count += int(i)
    }
  }

  return count, err
}

func (client HttpGithubClient) Organizations(login string) ([]string, error) {
  url := fmt.Sprintf("https://api.github.com/users/%s/orgs", login)
  body, err := client.Request(url)
  if err != nil {
    log.Fatalf("error requesting organizations for user %+v", login)
    return []string{}, err
  }
  orgResp := []OrgResponse {}
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
  Organization  string `json:"login"`
}

func NewGithubClient(wrappers ...wrapper) HttpGithubClient {
  return HttpGithubClient { wrappers: wrappers }
}

type User struct {
  Login        string
  Id           int
  Name         string
  Location     string
  Company      string
  Followers    int
  PublicRepos  int `json:"public_repos"`
}

type UserSearchQuery struct {
  q           string
  sort        string
  order       string
  per_page    int
  pages       int
}
