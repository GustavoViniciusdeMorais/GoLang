package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", check)
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

func check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "It is alive!")
}
