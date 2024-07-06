package repository

import (
	"example.com/internal/core/domain"
	"gorm.io/gorm"
)

type UserGormRepository struct {
	db *gorm.DB
}

func NewUserGormRepository(db *gorm.DB) *UserGormRepository {
	return &UserGormRepository{db: db}
}

func (r *UserGormRepository) Save(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *UserGormRepository) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	result := r.db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserGormRepository) FindAll() ([]*domain.User, error) {
	var users []*domain.User
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *UserGormRepository) Delete(id uint) error {
	return r.db.Delete(&domain.User{}, id).Error
}
