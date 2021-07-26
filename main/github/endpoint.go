package github

import (
	"encoding/json"
	"github.com/go-kit/kit/log/level"
	"github.com/gorilla/mux"
	"github_go/main/logger"
	"net/http"
)

type (
	FetchUserRepositoryResponse struct {
		Name     string `json:"name"`
		Owner    string `json:"owner"`
		Language string `json:"language"`
		Forks    int    `json:"forks"`
		Stars    int    `json:"stars"`
	}
)

func fetchGithubUserRepositoryInfoEndpoint(w http.ResponseWriter, r *http.Request) {
	level.Info(logger.Logger).Log("msg", "Endpoint hit: fetchGithubUserRepositoryInfoEndpoint")

	vars := mux.Vars(r)
	user := vars["user"]
	repository := vars["repository"]
	level.Info(logger.Logger).Log("msg", "user: "+user+", repository: "+repository)

	var githubRepository, err = fetchGithubRepo(user, repository)
	handleError(err, w)

	var response = FetchUserRepositoryResponse{
		githubRepository.Name,
		githubRepository.Owner.Login,
		githubRepository.Language,
		githubRepository.Forks,
		githubRepository.Stars,
	}

	err = json.NewEncoder(w).Encode(response)
	handleError(err, w)
}

func handleError(err error, w http.ResponseWriter) {
	if err != nil {
		level.Error(logger.Logger).Log("error", err)
		w.WriteHeader(http.StatusInternalServerError)
		panic(err)
	}
}
