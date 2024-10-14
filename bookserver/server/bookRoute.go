package server

import (
	"encoding/json"
	"log"
	"net/http"

	"xpJain.co/bookserver/models"
	"xpJain.co/bookserver/services"
)

var bookRouter = Router.PathPrefix("/books").Subrouter().StrictSlash(true)


// post /books
// get /books
// get /books/{id}
// put /books/{id}
// delete /books/{id}

func BookRouteInit() {


	services.InitBookDB();
	
	// Create a new router
	bookRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := services.GetBook()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}).Methods("GET")

	bookRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var book models.Book

		err := json.NewDecoder(r.Body).Decode(&book)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}		

		services.AddBook(book)
		log.Default().Println("Book added")

	}).Methods("POST")
}