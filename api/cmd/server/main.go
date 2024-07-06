package main

import (
	"log"
	"net/http"

	"example.com/internal/adapter/config"
	"example.com/internal/adapter/handler/myhttp"
	"example.com/internal/adapter/storage/postgres"
	"example.com/internal/adapter/storage/postgres/repository"
	"example.com/internal/core/domain"
	"example.com/internal/core/service"
)

func main() {

	config := config.NewDB()

	db, err := postgres.NewPostgresDB(*config)
	if err != nil {
		log.Fatal(err)
	}

	// Migrate the schema
	db.AutoMigrate(&domain.User{})

	userRepo := repository.NewUserGormRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := myhttp.NewUserHandler(userService)
	router := myhttp.NewRouter(userHandler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
