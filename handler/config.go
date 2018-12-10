package handler

import (
	"encoding/json"
	"net/http"
)

// Config defines config file structure
type Config struct {
	Fields   []Filed  `json:"fields"`
	Provider Provider `json:"provider"`

	Output struct {
		Format string `json:"format"`
		Path   string `json:"path"`
	} `json:"output"`
}

var configs map[string]Config

func init() {
	res, err := http.Get("https://raw.githubusercontent.com/sijad/serverlessman/master/now/config.json")
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	dec := json.NewDecoder(res.Body)

	dec.Decode(&configs)
}
