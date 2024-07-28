package service

import (
	os "os"
	"strconv"

	"example.com/internal/core/domain"
	"example.com/internal/core/port"
)

type UserService struct {
	repo port.UserRepository
}

func NewUserService(userRepo port.UserRepository) port.UserService {
	return &UserService{repo: userRepo}
}

func (s *UserService) CreateUser(name string, email string, birthday string, password string) (*domain.User, error) {
	qtyUsers, err := s.repo.Count()
	if err != nil {
		return nil, err
	}
	users_limit, err := strconv.ParseInt(os.Getenv("USERS_LIMIT"), 10, 64)
	if err != nil {
		panic(err)
	}
	if qtyUsers >= users_limit {
		return nil, nil
	}
	user := &domain.User{Name: name, Email: email, Birthday: birthday, Password: password, Active: false}
	user, err = s.repo.Save(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) FindByEmail(email string) (*domain.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *UserService) ListUsers(page string, limit string) ([]*domain.User, error) {
	pagination := domain.NewPagination(page, limit)
	pagination = pagination.CalculatePagination()
	return s.repo.FindAll(pagination)
}

func (s *UserService) DeleteUser(id uint) error {
	return s.repo.Delete(id)
}
