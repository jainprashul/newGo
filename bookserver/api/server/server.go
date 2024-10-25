package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"xpJain.co/bookserver/server/middleware"
)

// Function to create a new router
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	
	// Add Middleware
	router.Use(middleware.Logger)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	return router
}




var Router = NewRouter()