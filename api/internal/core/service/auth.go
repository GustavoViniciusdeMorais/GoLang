package service

import (
	"example.com/internal/core/domain"
	"example.com/internal/core/port"
)

type AuthService struct {
	userRepository port.UserRepository
}

func (s *AuthService) Login(email string, password string) (*domain.User, error) {
	user, err := s.userRepository.Login(email, password)
	if err != nil {
		return &domain.User{}, err
	}
	return user, nil
}

func (s *AuthService) FindByEmail(email string) (*domain.User, error) {
	user, err := s.userRepository.FindByEmail(email)
	if err != nil {
		return &domain.User{}, err
	}
	return user, nil
}

func NewAuthService(userRepository port.UserRepository) port.AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}
