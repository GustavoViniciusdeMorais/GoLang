package main

import (
	"context"
	"fmt"
	"log"

	"example.com/internal/adapter/config"
	"example.com/internal/adapter/handler/myhttp"
	"example.com/internal/adapter/storage/postgres"
	"example.com/internal/adapter/storage/postgres/repository"
	"example.com/internal/adapter/storage/redis"
	"example.com/internal/core/domain"
	"example.com/internal/core/service"
)

func main() {

	ctx := context.Background()

	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	postgresDB, err := postgres.NewPostgresDB(*config.DB)
	if err != nil {
		log.Fatal(err)
	}

	cache, err := redis.New(ctx, config.Redis)
	if err != nil {
		log.Fatal(err)
	}
	defer cache.Close()

	// Migrate the schema
	postgresDB.DB.AutoMigrate(&domain.User{})

	userRepo := repository.NewUserGormRepository(postgresDB.DB)
	userService := service.NewUserService(userRepo)
	userHandler := myhttp.NewUserHandler(userService)

	authService := service.NewAuthService(userRepo)
	authHandler := myhttp.NewAuthHandler(authService, cache)

	server := myhttp.NewServer()
	err = server.RegisterRoutes(
		cache,
		userHandler,
		authHandler,
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := server.Start(fmt.Sprintf("0.0.0.0:%s", config.HTTP.Port)); err != nil {
		log.Fatal(err)
	}
}
