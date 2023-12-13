package api

import (
	"fmt"
	"log"
	"net/http"

	"example.com/cmds"
	"example.com/handlers"
	"github.com/gorilla/mux"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	users := cmds.All()
	fmt.Fprint(w, users)
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	response := handlers.Register()
	fmt.Fprint(w, response)
}

func Run() {

	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/register", registerUser).Methods("POST")

	http.Handle("/", r)
	fmt.Println("Server at localhost:9003")
	log.Fatal(http.ListenAndServe(":9003", nil))
}
