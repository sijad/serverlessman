package main

import (
	"net/http"
)

// Handler will be executed by now
func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/submit":
		submit(w, r)
	default:
		w.WriteHeader(404)
	}
	// fmt.Fprintf(w, r.URL.Path)
}
