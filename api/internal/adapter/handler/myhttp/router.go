package myhttp

import (
	"example.com/internal/adapter/config"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Engine *gin.Engine
}

func NewRouter(
	config *config.HTTP,
	handler *UserHandler,
) (*Route, error) {
	router := gin.New()

	v1 := router.Group("/api/v1")
	{
		user := v1.Group("/users")
		{
			user.GET("", handler.GetUsers)
		}
	}

	return &Route{
		Engine: router,
	}, nil
}

func (r *Route) Start(listenAddr string) error {
	return r.Engine.Run(listenAddr)
}
