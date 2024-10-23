package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Function to create a new router
func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	return router
}

// Function to create a new route
func NewRoute(router *mux.Router, path string, handler func(http.ResponseWriter, *http.Request)) {
	router.HandleFunc(path, handler)
}




var Router = NewRouter()