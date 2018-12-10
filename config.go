package main

import (
	"encoding/json"
	"os"
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
	jsonFile, err := os.Open("./config.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	dec := json.NewDecoder(jsonFile)

	dec.Decode(&configs)
}
