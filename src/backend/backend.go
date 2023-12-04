package backend

import (
	"fmt"
	"log"
	"net/http"
)

func myEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Test endpoint\n")
}

func Run(addr string) {
	http.HandleFunc("/", myEndpoint)
	fmt.Println("Server at port", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
