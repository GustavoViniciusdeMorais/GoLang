package port

import "example.com/internal/core/domain"

type UserRepository interface {
	Save(user *domain.User) error
	FindByEmail(email string) (*domain.User, error)
	FindAll(pagination *domain.Pagination) ([]*domain.User, error)
	Delete(id uint) error
	Login(email string, password string) (*domain.User, error)
}

type UserService interface {
	CreateUser(name, email string) error
	FindByEmail(email string) (*domain.User, error)
	ListUsers(page string, limit string) ([]*domain.User, error)
	DeleteUser(id uint) error
}
