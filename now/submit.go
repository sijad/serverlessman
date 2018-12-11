package main

import (
	"net/http"
	"os"

	"github.com/sijad/serverlessman/handler"
)

// Handler will be executed by now
func Handler(w http.ResponseWriter, r *http.Request) {
	handler.Submit(w, r)
}

func init() {
	token := os.Getenv("GITHUB_TOKEN")
	repo := os.Getenv("GITHUB_REPO")
	owner := os.Getenv("GITHUB_OWNER")
	branch := os.Getenv("GITHUB_BRANCH")

	handler.InitProvider(token, repo, owner, branch)
	handler.InitConfigs()
}
