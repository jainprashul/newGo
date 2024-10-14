package services

import (
	"fmt"

	"xpJain.co/bookserver/db"
	"xpJain.co/bookserver/models"
)

var bookDB, err = db.GetDB("books")


func InitBookDB() {
	// Create a new DB if it does not exist
	if err != nil {
		bookDB = db.CreateDB("books")
	}
}




func GetBook() []models.Book {
	var books []models.Book
	err := bookDB.GetAll(&books)
	if err != nil {
		fmt.Println(err)
	}
	return books
}

func AddBook(book models.Book) {

	book.ID = fmt.Sprintf("book-%d", len(GetBook())+1)

	books := GetBook()
	books = append(books, book)
	bookDB.AddData(books)
}

func DeleteBook(book models.Book) {
	b := GetBook()
	for i, v := range b {
		if v == book {
			b = append(b[:i], b[i+1:]...)
		}
	}

	bookDB.AddData(b)
}
