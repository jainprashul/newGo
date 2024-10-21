package main

import (
	"fmt"
	"net/http"

	"xpJain.co/bookserver/server"
)



func main() {
	// Create a new router
	router := server.Router

	// Initialize the book route
	server.BookRouteInit();


	// Start the server
	fmt.Println(`Server is running on port 8000`)
	http.ListenAndServe(":8000", router)
}