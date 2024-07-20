package myhttp

import (
	"context"

	"example.com/internal/core/domain"
	"example.com/internal/core/port"
)

type AuthHandler struct {
}

func (a *AuthHandler) Login(ctx context.Context, email string, password string) (string, error) {
	return "", nil
}
func (a *AuthHandler) CreateToken(user *domain.User) (string, error) {
	return "", nil
}
func (a *AuthHandler) VerifyToken(token string) (string, error) {
	return "", nil
}

func NewAuthHandler() port.AuthService {
	return &AuthHandler{}
}
