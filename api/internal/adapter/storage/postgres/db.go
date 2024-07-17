package postgres

import (
	"fmt"

	"example.com/internal/adapter/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	DB *gorm.DB
}

func NewPostgresDB(cfg config.DB) (PostgresDB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return PostgresDB{DB: nil}, err
	}

	postgresDb := PostgresDB{DB: db}

	return postgresDb, nil
}
