package http

import (
	"github.com/gorilla/mux"
)

func NewRouter(handler *UserHandler) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/users", handler.GetUsers).Methods("GET")
	return router
}
