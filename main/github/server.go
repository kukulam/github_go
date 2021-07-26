package github

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var httpClient http.Client

func init() {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 20
	t.MaxConnsPerHost = 20
	t.MaxIdleConnsPerHost = 20

	httpClient = http.Client{
		Timeout:   time.Duration(1) * time.Second,
		Transport: t,
	}
}

func NewHTTPRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/github/users/{user}/{repository}", fetchGithubUserRepositoryInfoEndpoint)
	return router
}
