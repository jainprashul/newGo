package server

import (
	"xpJain.co/bookserver/db"
	"xpJain.co/bookserver/models"
)


func BookRouteInit() {
	// Create a new BookDB
	bookDB := db.CreateDB("books")
	bookService := NewRouteService[*models.Book](bookDB)

	bookService.InitService()
}
