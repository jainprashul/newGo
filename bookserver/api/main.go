package main

import (
	"fmt"
	"net/http"
	"os"

	"xpJain.co/bookserver/db"
	"xpJain.co/bookserver/server"
)



func main() {

	// Initialize the database
	db.InitializeDB()

	// Create a new router
	router := server.Router

	// Initialize the book route
	server.BookRouteInitize()

	// GET environment variable PORT 
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8000"
	}

	// Start the server
	fmt.Println(`Server is running on port ` + PORT)
	http.ListenAndServe(":"+PORT, router)
}