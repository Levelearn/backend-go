package repository

import (
	"levelearn-backend/internal/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsername(username string) (*entity.User, error)
	Create(user *entity.User) error
	FindById(id int) (*entity.User, error)
	UpdateProfile(id int, name *string, image *string) error
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

func (r *userRepository) FindById(id int) (*entity.User, error) {
	var user entity.User

	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepository) UpdateProfile(
	id int,
	name *string,
	image *string,
) error {
	updates := map[string]interface{}{}

	if name != nil {
		updates["name"] = *name
	}

	if image != nil {
		updates["image"] = *image
	}

	if len(updates) == 0 {
		return nil
	}

	return r.db.
		Model(&entity.User{}).
		Where("id = ?", id).
		Updates(updates).
		Error
}
