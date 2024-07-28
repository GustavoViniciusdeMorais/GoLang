package myhttp

import (
	"log"
	"net/http"

	"example.com/internal/core/domain"
	"example.com/internal/core/port"
	"github.com/labstack/echo/v4"
)

type EchoServer struct {
	echo *echo.Echo
}

func NewServer() *EchoServer {
	server := &EchoServer{
		echo: echo.New(),
	}
	return server
}

func (s *EchoServer) RegisterRoutes(
	redisCache port.CacheRepository,
	userHandler *UserHandler,
	authHandler *AuthHandler,
) error {
	s.echo.GET("/liveness", s.Liveness)

	ag := s.echo.Group("/auth")
	ag.POST("/login", authHandler.Login)
	ag.POST("/logout", authHandler.Logout)

	ug := s.echo.Group("/users")
	ug.Use(JWTMiddleware(redisCache))
	ug.GET("", userHandler.GetUsers)
	ug.POST("", userHandler.CreateUser)

	return nil
}

func (s *EchoServer) Start(listenAddr string) error {
	if err := s.echo.Start(listenAddr); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server shutdown occurred: %s", err)
		return err
	}
	return nil
}

func (s *EchoServer) Liveness(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, domain.Health{Status: "OK"})
}
