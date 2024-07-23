package service

import (
	"example.com/internal/core/domain"
	"example.com/internal/core/port"
)

type UserService struct {
	repo port.UserRepository
}

func NewUserService(userRepo port.UserRepository) port.UserService {
	return &UserService{repo: userRepo}
}

func (s *UserService) CreateUser(name, email string) error {
	user := &domain.User{Name: name, Email: email}
	return s.repo.Save(user)
}

func (s *UserService) FindByEmail(email string) (*domain.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *UserService) ListUsers() ([]*domain.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
