package repository

import (
	"levelearn-backend/internal/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsername(username string) (*entity.User, error)
	Create(user *entity.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) FindByUsername(username string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *userRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}
