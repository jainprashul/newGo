package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"xpJain.co/bookserver/db"
)

// GORM_ROUTE_HANDLER is a struct that contains the db model instance
type GORM_ROUTE_HANDLER[T Indentifiable] struct{
	db *db.DBModel[T]
	router *mux.Router
	
}

// NewRouteService is a function that creates a new RouteService struct with the provided db 
func New_GormRouteHandler[T Indentifiable](db *db.DBModel[T]) *GORM_ROUTE_HANDLER[T] {
	return &GORM_ROUTE_HANDLER[T]{
		db: db,
		router: Router.PathPrefix("/" + db.GetTableName()).Subrouter().StrictSlash(true),
	}
}

func (gr *GORM_ROUTE_HANDLER[T]) GetRouter() *mux.Router {
	return gr.router
}


func (gr *GORM_ROUTE_HANDLER[T]) InitService() {
	subroute := gr.router

	subroute.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		objects, err := gr.db.GetObjects()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&objects)
	}).Methods("GET")

	subroute.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		object, err := gr.db.Get(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(object)
	}).Methods("GET")

	subroute.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var object T
		err := json.NewDecoder(r.Body).Decode(&object)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = gr.db.Create(object)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// 201 Created
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(object)
	}).Methods("POST")

	subroute.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		
		var object T
		err := json.NewDecoder(r.Body).Decode(&object)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = gr.db.Update(object , id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// 204 No Content
		w.WriteHeader(http.StatusNoContent)
	}).Methods("PUT")

	subroute.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		err := gr.db.Delete(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// 204 No Content
		w.WriteHeader(http.StatusNoContent)
	}).Methods("DELETE")
}


