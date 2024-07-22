package port

import (
	"example.com/internal/core/domain"
)

type AuthService interface {
	Login(email string, password string) (*domain.User, error)
}
