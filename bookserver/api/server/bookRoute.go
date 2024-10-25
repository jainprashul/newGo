package server

import (
	"xpJain.co/bookserver/db"
	"xpJain.co/bookserver/models"
	"xpJain.co/bookserver/server/middleware"
)

func BookRouteInit() {
	// Create a new BookDB
	bookDB := db.CreateFileDB("books")
	bookService := NewRouteService[*models.Book](bookDB)

	bookService.InitService()
}


func BookRouteInitize() {
	// Create a new BookDB
	bookDB := db.NewModel("books", &models.Book{})

	route := New_GormRouteHandler(bookDB)

	route.router.Use(middleware.Authentication)

	route.InitService()

}