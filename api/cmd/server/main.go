package main

import (
	"fmt"
	"log"

	"example.com/internal/adapter/config"
	"example.com/internal/adapter/handler/myhttp"
	"example.com/internal/adapter/storage/postgres"
	"example.com/internal/adapter/storage/postgres/repository"
	"example.com/internal/core/domain"
	"example.com/internal/core/service"
)

func main() {

	configDB := config.NewDB()
	congigHost := config.NewHTTP()

	postgresDB, err := postgres.NewPostgresDB(*configDB)
	if err != nil {
		log.Fatal(err)
	}

	// Migrate the schema
	postgresDB.DB.AutoMigrate(&domain.User{})

	userRepo := repository.NewUserGormRepository(postgresDB.DB)
	userService := service.NewUserService(userRepo)
	userHandler := myhttp.NewUserHandler(userService)

	server := myhttp.NewServer()
	if err := server.RegisterRoutes(userHandler); err != nil {
		log.Fatal(err)
	}
	if err := server.Start(fmt.Sprintf("0.0.0.0:%s", congigHost.Port)); err != nil {
		log.Fatal(err)
	}
}
