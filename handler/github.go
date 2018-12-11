package handler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

const (
	githubURI = "https://api.github.com/repos/%s/%s"
)

type github struct {
	Token  string
	Branch string
	Repo   string
	Owner  string
}

type githubContent struct {
	Content string `json:"content"`
}

type githubHeadRef struct {
	Object struct {
		Sha string `json:"sha"`
	} `json:"object"`
}

type githubNewRef struct {
	Ref string `json:"ref"`
	Sha string `json:"sha"`
}

type githubNewPR struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Head  string `json:"head"`
	Base  string `json:"base"`
}

type githubNewContent struct {
	Message string `json:"message"`
	Branch  string `json:"branch"`
	Content string `json:"content"`
}

func (g *github) url() string {
	return fmt.Sprintf(githubURI, g.Owner, g.Repo)
}

func (g *github) get(path string, result interface{}) error {
	res, err := getJSON(g.url()+path, g.headers())

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New("github get failed with status: " + res.Status)
	}

	if err := json.NewDecoder(res.Body).Decode(result); err != nil {
		return err
	}

	return nil
}

func (g *github) put(path string, body interface{}) error {
	res, err := sendJSON("PUT", g.url()+path, g.headers(), body)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	return nil
}

func (g *github) post(path string, body interface{}) error {
	res, err := sendJSON("POST", g.url()+path, g.headers(), body)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	return nil
}

func (g *github) headers() map[string]string {
	headers := make(map[string]string)
	headers["Authorization"] = "token " + g.Token
	headers["Content-Type"] = "application/json"
	return headers
}

func (g *github) CreateFile(path string, body io.Reader, branch string) error {
	b := new(bytes.Buffer)
	encoder := base64.NewEncoder(base64.StdEncoding, b)
	io.Copy(encoder, body)

	newContent := &githubNewContent{
		Branch:  branch,
		Message: "new file",
		Content: b.String(),
	}

	return g.put("/contents/"+path, newContent)
}

func (g *github) CreateBranch(name string) error {
	var gr githubHeadRef
	if err := g.get("/git/refs/heads/"+g.Branch, &gr); err != nil {
		return err
	}

	newRef := &githubNewRef{
		Ref: "refs/heads/" + name,
		Sha: gr.Object.Sha,
	}

	if err := g.post("/git/refs", newRef); err != nil {
		return err
	}

	return nil
}

func (g *github) CreateModerationRequest(branch string) error {
	newPR := &githubNewPR{
		Title: "new entry",
		Body:  "TODO",
		Head:  branch,
		Base:  g.Branch,
	}

	if err := g.post("/pulls", newPR); err != nil {
		return err
	}

	return nil
}

// CreateNewFile Create a new file in git repo
func (g *github) CreateNewFile(path string, body io.Reader, moderation bool) error {
	branch := g.Branch

	if moderation {
		uid, err := uuid.NewV4()

		if err != nil {
			return err
		}

		branch = "serverlessman-" + uid.String()

		if err := g.CreateBranch(branch); err != nil {
			return err
		}
	}

	if err := g.CreateFile(path, body, branch); err != nil {
		return err
	}

	if moderation {
		g.CreateModerationRequest(branch)
	}

	return nil
}

// GetRepoConfigs Create a new file in git repo
func (g *github) GetRepoConfigs() (map[string]Config, error) {
	var content githubContent
	if err := g.get("/contents/serverlessman.json?ref="+g.Branch, &content); err != nil {
		return nil, err
	}

	c, err := base64.URLEncoding.DecodeString(content.Content)
	if err != nil {
		return nil, err
	}

	var configs map[string]Config
	if err := json.Unmarshal(c, &configs); err != nil {
		return nil, err
	}

	return configs, nil
}
