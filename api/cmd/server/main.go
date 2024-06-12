package main

import (
	"log"
	"net/http"

	"github.com/GustavoViniciusdeMorais/api/internal/adapters/http"
	"github.com/GustavoViniciusdeMorais/api/internal/core/service"
)

func main() {
	userService := service.NewUserService()
	userHandler := http.NewUserHandler(userService)
	router := http.NewRouter(userHandler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
