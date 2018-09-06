package net

import "net/http"
import "fmt"
import "io/ioutil"

type Requester func(req *http.Request) ([]byte, error)

type Wrapper func(Requester) Requester

func Compose(wrappers ...Wrapper) Wrapper {
	return func(r Requester) Requester {
		for _, wrapper := range wrappers {
			r = wrapper(r)
		}
		return r
	}
}

func TokenAuth(token string) Wrapper {
	return func(r Requester) Requester {
		return func(req *http.Request) ([]byte, error) {
			req.Header.Add("Authorization", fmt.Sprintf("bearer %s", token))
			return r(req)
		}
	}
}

func MakeRequester(client *http.Client) Requester {
	return func(req *http.Request) ([]byte, error) {
		resp, err := client.Do(req)

		if err != nil {
			return []byte{}, err
		}

		bodyText, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return []byte{}, err
		}

		return bodyText, nil
	}
}
