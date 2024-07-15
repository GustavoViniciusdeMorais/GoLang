package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"example.com/internal/adapter/config"
	"example.com/internal/adapter/handler/myhttp"
	"example.com/internal/adapter/storage/postgres"
	"example.com/internal/adapter/storage/postgres/repository"
	"example.com/internal/core/domain"
	"example.com/internal/core/service"
)

func main() {

	configDB := config.NewDB()
	congiHost := config.NewHTTP()

	db, err := postgres.NewPostgresDB(*configDB)
	if err != nil {
		log.Fatal(err)
	}

	// Migrate the schema
	db.AutoMigrate(&domain.User{})

	userRepo := repository.NewUserGormRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := myhttp.NewUserHandler(userService)

	listenAddr := fmt.Sprintf("%s:%s", config.HTTP.Url, config.HTTP.Port)
	router := myhttp.NewRouter(userHandler)

	err = router.Serve(listenAddr)
	if err != nil {
		os.Exit(1)
	}

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
