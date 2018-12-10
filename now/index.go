package main

import (
	"github.com/sijad/serverlessman/handler"
	"net/http"
)

// Handler will be executed by now
func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/submit":
		handler.Submit(w, r)
	default:
		w.WriteHeader(404)
	}
	// fmt.Fprintf(w, r.URL.Path)
}
