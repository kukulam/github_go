package main

import (
	"fmt"
	"github.com/go-kit/kit/log/level"
	"github_go/main/github"
	"github_go/main/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	errs := make(chan error)

	level.Info(logger.Logger).Log("msg", "Starting app")
	defer level.Info(logger.Logger).Log("msg", "Stopping app")

	// HTTP server
	go func() {
		port := ":8080"
		level.Info(logger.Logger).Log("msg", "Listening on port"+port)
		router := github.NewHTTPRouter()
		errs <- http.ListenAndServe(port, router)
	}()

	// CTRL+C
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	level.Error(logger.Logger).Log("exit", <-errs)
}
