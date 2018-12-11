package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sijad/serverlessman/handler"
)

// consts
const (
	PORT = "8084"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("GITHUB_TOKEN")
	repo := os.Getenv("GITHUB_REPO")
	owner := os.Getenv("GITHUB_OWNER")
	branch := os.Getenv("GITHUB_BRANCH")
	handler.InitProvider(token, repo, owner, branch)
	handler.InitConfigs()

	http.HandleFunc("/", handler.Submit)
	log.Printf("connect to http://localhost:%s/ for access\n", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}
