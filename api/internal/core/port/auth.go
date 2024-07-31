package port

import (
	"example.com/internal/core/domain"
)

type AuthService interface {
	FindByEmail(email string) (*domain.User, error)
}
