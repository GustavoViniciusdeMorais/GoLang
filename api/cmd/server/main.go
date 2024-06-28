package main

import (
	"log"
	"net/http"

	"example.com/internal/adapter/handler/myhttp"
	"example.com/internal/core/service"
)

func main() {
	userService := service.NewUserService()
	userHandler := myhttp.NewUserHandler(userService)
	router := myhttp.NewRouter(userHandler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
