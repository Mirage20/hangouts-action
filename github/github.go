package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GitHub struct {
	*http.Client
	token string
	repo  string
}

func NewClient(token string, repo string) *GitHub {
	return &GitHub{
		Client: &http.Client{},
		token:  token,
		repo:   repo,
	}
}

func (g *GitHub) CheckRuns(sha string) (*CheckRunsResponse, error) {
	b, err := g.doGet(fmt.Sprintf("commits/%s/check-runs", sha))
	if err != nil {
		return nil, err
	}
	resp := &CheckRunsResponse{}
	err = json.Unmarshal(b, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GitHub) Status(sha string) (*StatusResponse, error) {
	b, err := g.doGet(fmt.Sprintf("commits/%s/status", sha))
	if err != nil {
		return nil, err
	}
	resp := &StatusResponse{}
	err = json.Unmarshal(b, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (g *GitHub) Validate() ([]byte, error) {
	return g.doGet("")
}

func (g *GitHub) doGet(api string) ([]byte, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s", g.repo)
	if len(api) > 0 {
		url = fmt.Sprintf("https://api.github.com/repos/%s/%s", g.repo, api)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "token "+g.token)
	req.Header.Add("Accept", "application/vnd.github.v3+json; application/vnd.github.antiope-preview+json")
	if err != nil {
		log.Fatal(err)
	}
	res, err := g.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode/100 != 2 {
		body, _ := ioutil.ReadAll(res.Body)
		return nil, fmt.Errorf("github error: %s", body)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
