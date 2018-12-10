package handler

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
)

const (
	githubURI = "https://api.github.com/repos/%s/%s"
)

type github struct {
	token  string
	branch string
	repo   string
	owner  string
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
	return fmt.Sprintf(githubURI, g.owner, g.repo)
}

func (g *github) get(path string, result interface{}) error {
	res, err := getJSON(g.url()+path, g.headers())

	if err != nil {
		return err
	}

	defer res.Body.Close()

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
	headers["Authorization"] = "token " + g.token
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
	if err := g.get("/git/refs/heads/"+g.branch, &gr); err != nil {
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
		Base:  g.branch,
	}

	if err := g.post("/pulls", newPR); err != nil {
		return err
	}

	return nil
}
