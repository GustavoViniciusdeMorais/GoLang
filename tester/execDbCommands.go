package main

import (
	"log"

	"example.com/backend"
)

func main() {
	db, err := backend.DbConnection()
	if err != nil {
		log.Fatal(err)
	}
	backend.ListProducts(db)
}
