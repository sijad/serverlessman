package handler

import (
	"io"
)

var provider Provider

// Provider defines provider structure
type Provider interface {
	CreateBranch(name string) error
	CreateFile(path string, body io.Reader, branch string) error
	CreateModerationRequest(branch string) error
	CreateNewFile(path string, body io.Reader, moderation bool) error
	GetRepoConfigs() (map[string]Config, error)
}

// InitProvider init provider and configs
func InitProvider(token, repo, owner, branch string) {
	if token == "" {
		panic("github token must be provided")
	}

	if repo == "" {
		panic("github repo must be provided")
	}

	if repo == "" {
		panic("github repo must be provided")
	}

	if repo == "" {
		panic("github repo must be provided")
	}

	provider = &github{
		Token:  token,
		Branch: branch,
		Repo:   repo,
		Owner:  owner,
	}
}
