package config

import "os"

type (
	Container struct {
		DB    *DB
		HTTP  *HTTP
		Redis *Redis
	}

	DB struct {
		Host     string
		User     string
		Password string
		Database string
		Port     string
		SSLMode  string
	}

	HTTP struct {
		Url  string
		Port string
	}

	Redis struct {
		Host     string
		Password string
		Port     string
	}
)

func New() (*Container, error) {

	db := &DB{
		Host:     os.Getenv("POSTGRES_HOST"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
		Port:     os.Getenv("POSTGRES_PORT"),
		SSLMode:  "disable",
	}

	http := &HTTP{
		Url:  os.Getenv("HTTP_URL"),
		Port: os.Getenv("HTTP_PORT"),
	}

	redis := &Redis{
		Host:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		Port:     os.Getenv("REDIS_PORT"),
	}

	return &Container{
		DB:    db,
		HTTP:  http,
		Redis: redis,
	}, nil
}
