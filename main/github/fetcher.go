package github

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type (
	FetchUserGithubRepositoryResponse struct {
		Name     string      `json:"name"`
		Owner    GithubOwner `json:"owner"`
		Language string      `json:"language"`
		Forks    int         `json:"forks"`
		Stars    int         `json:"stars"`
	}

	GithubOwner struct {
		Login string `json:"login"`
	}
)

func fetchGithubRepo(user string, repository string) (*FetchUserGithubRepositoryResponse, error) {
	resp, err := httpClient.Get("https://api.github.com/reposS/" + user + "/" + repository)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic("problem with closing buffer")
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var githubRepositoryBody FetchUserGithubRepositoryResponse
	err = json.Unmarshal(body, &githubRepositoryBody)
	if err != nil {
		return nil, err
	}

	return &githubRepositoryBody, err
}
