package backend

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "GET Request\n")
}

func postRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "POST Request\n")
}

func deleteRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "DELETE Request\n")
}

func Run(addr string) {
	r := mux.NewRouter()
	r.HandleFunc("/products", getRequest).Methods("GET")
	r.HandleFunc("/products", postRequest).Methods("POST")
	r.HandleFunc("/products", deleteRequest).Methods("DELETE")

	http.Handle("/products", r)
	fmt.Println("Server at localhost:9003")
	log.Fatal(http.ListenAndServe(":9003", nil))
}
