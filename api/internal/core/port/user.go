package port

import "example.com/internal/core/domain"

type UserRepository interface {
	Save(user *domain.User) error
	FindByID(id uint) (*domain.User, error)
	FindAll() ([]*domain.User, error)
	Delete(id uint) error
}

type UserService interface {
	CreateUser(name, email string) error
	GetUserByID(id uint) (*domain.User, error)
	ListUsers() ([]*domain.User, error)
	DeleteUser(id uint) error
}
