package config

import "os"

type DB struct {
	Host     string
	User     string
	Password string
	Database string
	Port     string
	SSLMode  string
}

func NewDB() *DB {
	return &DB{
		Host:     os.Getenv("POSTGRES_HOST"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
		Port:     os.Getenv("POSTGRES_PORT"),
		SSLMode:  "disable",
	}
}

type HTTP struct {
	Url  string
	Port string
}

func NewHTTP() *HTTP {
	return &HTTP{
		Url:  os.Getenv("HTTP_URL"),
		Port: os.Getenv("HTTP_PORT"),
	}
}
