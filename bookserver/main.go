package main

import (
	"net/http"

	"xpJain.co/bookserver/server"
)



func main() {
	// Create a new router
	router := server.Router

	// Initialize the book route
	server.BookRouteInit();


	http.ListenAndServe(":8000", router)
}