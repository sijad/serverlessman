package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// consts
const (
	PORT = "8084"
)

func main() {
	initEnv()
	http.HandleFunc("/", Handler)
	log.Printf("connect to http://localhost:%s/ for access\n", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

type nowJSON struct {
	Env map[string]string `json:"env"`
}

func initEnv() {
	jsonFile, err := os.Open("./now.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	dec := json.NewDecoder(jsonFile)

	var res nowJSON
	dec.Decode(&res)

	for k, v := range res.Env {
		if v[0] == '@' {
			if os.Getenv(k) == "" {
				panic("can not find secure env " + k)
			}
			continue
		}

		os.Setenv(k, v)
	}
}
