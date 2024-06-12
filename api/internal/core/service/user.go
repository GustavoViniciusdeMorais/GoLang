package service

import "github.com/GustavoViniciusdeMorais/api/internal/core/domain"

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUsers() []domain.User {
	// Mockup data
	return []domain.User{
		{ID: 1, Name: "John Doe", Age: 30},
		{ID: 2, Name: "Jane Doe", Age: 25},
	}
}
