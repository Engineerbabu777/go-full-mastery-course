package router

import "github.com/gorilla/mux"

func Router() *mux.Router {

	r := mux.NewRouter()

	// ROUTES!
	// r.HandleFunc("/api/create-todo", ).Methods("POST");

	return r
}
