package main

import (
	"log"

	"example.com/backend"
	"example.com/practice"
)

func main() {
	practice.Test()
	// backend.Run("localhost:9003")
	db, err := backend.DbConnection()
	if err != nil {
		log.Fatal(err)
	}
	backend.CreateTableProducts(db)
	backend.CreateSampleData(db)
}
