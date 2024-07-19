package port

import (
	"context"

	"example.com/internal/core/domain"
)

type AuthService interface {
	Login(ctx context.Context, email string, password string) (string, error)
	CreateToken(user *domain.User) (string, error)
	VerifyToken(token string) (string, error)
}
