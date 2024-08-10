package port

import "example.com/internal/core/domain"

type UserRepository interface {
	Save(user *domain.User) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindById(id int64) (*domain.User, error)
	FindAll(pagination *domain.Pagination) ([]*domain.User, error)
	Update(user *domain.User) (*domain.User, error)
	Delete(id int64) error
	Count() (int64, error)
}

type UserService interface {
	Save(name string, email string, birthday string, password string) (*domain.User, error)
	FindByEmail(email string) (*domain.User, error)
	FindAll(page string, limit string) ([]*domain.User, error)
	FindById(id int64) (*domain.User, error)
	Update(id int, name string, email string, birthday string, password string, active bool) (*domain.User, error)
	Delete(id int64) error
}
