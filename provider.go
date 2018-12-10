package main

import (
	"io"
	"os"

	"github.com/satori/go.uuid"
)

// ProviderType type
type ProviderType string

// Git Providers
const (
	ProviderGithub ProviderType = "github"
)

// Provider defines provider structure
type Provider struct {
	Name ProviderType `json:"name"`

	Owner  string `json:"owner"`
	Repo   string `json:"repo"`
	Branch string `json:"branch"`

	Moderation bool `json:"moderation"`
}

func (p *Provider) provider() gitProvider {
	// TODO it might be useful to cache the result

	if p.Name == ProviderGithub {
		return &github{
			token:  os.Getenv("GITHUB_TOKEN"),
			branch: p.Branch,
			repo:   p.Repo,
			owner:  p.Owner,
		}
	}

	panic("provider not found")
}

// CreateNewFile Create a new file in git repo
func (p *Provider) CreateNewFile(path string, body io.Reader) error {
	gp := p.provider()
	branch := p.Branch

	if p.Moderation {
		uid, err := uuid.NewV4()

		if err != nil {
			return err
		}

		branch = "serverlessman-" + uid.String()

		if err := gp.CreateBranch(branch); err != nil {
			return err
		}
	}

	if err := gp.CreateFile(path, body, branch); err != nil {
		return err
	}

	if p.Moderation {
		gp.CreateModerationRequest(branch)
	}

	return nil
}

type gitProvider interface {
	CreateBranch(name string) error
	CreateFile(path string, body io.Reader, branch string) error
	CreateModerationRequest(branch string) error
}
