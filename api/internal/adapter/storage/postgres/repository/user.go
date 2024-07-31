package repository

import (
	"example.com/internal/core/domain"
	"example.com/internal/core/port"
	"gorm.io/gorm"
)

type UserGormRepository struct {
	db *gorm.DB
}

func NewUserGormRepository(db *gorm.DB) port.UserRepository {
	return &UserGormRepository{db: db}
}

func (r *UserGormRepository) Save(user *domain.User) (*domain.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserGormRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserGormRepository) FindAll(pagination *domain.Pagination) ([]*domain.User, error) {
	var users []*domain.User
	result := r.db.Model(&domain.User{}).
		Select("name, email, birthday, active").
		Offset(pagination.Offset).
		Limit(pagination.LimitInt).
		Scan(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *UserGormRepository) Delete(id uint) error {
	return r.db.Delete(&domain.User{}, id).Error
}

func (r *UserGormRepository) Login(email string, password string) (*domain.User, error) {
	var user domain.User
	result := r.db.Where("email = ? AND password = ?", email, password).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserGormRepository) Count() (int64, error) {
	var count int64
	if err := r.db.Model(&domain.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
