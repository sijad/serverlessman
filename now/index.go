package main

import (
	"github.com/sijad/serverlessman/handler"
	"net/http"
)

// Handler will be executed by now
func Handler(w http.ResponseWriter, r *http.Request) {
	handler.Submit(w, r)
}
