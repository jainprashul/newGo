package server

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"xpJain.co/bookserver/db"
)

type Indentifiable interface {
	GetID() string
	SetID(string)
}

type RouteService[T Indentifiable] struct {
	// RouteService
	db db.FileDB
}

func NewRouteService[C Indentifiable](db db.FileDB) *RouteService[C] {
	// NewRouteService
	return &RouteService[C]{
		db: db,
	}
}

func (r *RouteService[C]) getAll(w http.ResponseWriter, _ *http.Request) {
	var data []C
	err := r.db.GetAll(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (r *RouteService[C]) get(w http.ResponseWriter, req *http.Request) {
	var data []C
	id := mux.Vars(req)["id"]

	if id == "" {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}

	err := r.db.GetAll(&data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, d := range data {
		if d.GetID() == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(d)
			return
		}
	}

	http.Error(w, "Data not found", http.StatusNotFound)
}

func (r *RouteService[C]) add(w http.ResponseWriter, req *http.Request) {
	var data C

	err := json.NewDecoder(req.Body).Decode(&data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data.SetID(r.db.Name + "-" + uuid.NewString())

	var dataList []C
	err = r.db.GetAll(&dataList)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dataList = append(dataList, data)

	err = r.db.AddData(dataList)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func (r *RouteService[C]) delete(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	if id == "" {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}

	var data []C

	err := r.db.GetAll(&data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for i, d := range data {
		if d.GetID() == id {
			data = append(data[:i], data[i+1:]...)
			break
		}
	}

	err = r.db.AddData(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (r *RouteService[C]) update(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	if id == "" {
		http.Error(w, "ID not provided", http.StatusBadRequest)
		return
	}

	var data []C

	err := r.db.GetAll(&data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	print(id)

	for i, d := range data {
		if d.GetID() == id {
			var newData C

			err = json.NewDecoder(req.Body).Decode(&newData)
			newData.SetID(id)

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			data[i] = newData
			break
		}
	}

	err = r.db.AddData(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (r *RouteService[C]) GetDB() *db.FileDB {
	r.db.InitDB()
	return &r.db
}

func (r *RouteService[C]) GetRouter() *mux.Router {
	path := "/" + r.db.Name
	return Router.PathPrefix(path).Subrouter().StrictSlash(true)
}

func (r *RouteService[C]) InitService() {
	r.db.InitDB()

	// Create a new router for the service and add the routes to it
	path := "/" + r.db.Name
	subRoute := Router.PathPrefix(path).Subrouter().StrictSlash(true)

	subRoute.HandleFunc("/", r.getAll).Methods("GET")
	subRoute.HandleFunc("/", r.add).Methods("POST")
	subRoute.HandleFunc("/{id}", r.get).Methods("GET")
	subRoute.HandleFunc("/{id}", r.delete).Methods("DELETE")
	subRoute.HandleFunc("/{id}", r.update).Methods("PUT")

}