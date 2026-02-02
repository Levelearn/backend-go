package service

import (
	"errors"
	"levelearn-backend/internal/entity"
	"levelearn-backend/internal/repository"
)

type UserService interface {
	GetById(id int) (*entity.User, error)
	UpdateProfile(id int, name *string, image *string) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) GetById(id int) (*entity.User, error) {
	user, err := s.userRepo.FindById(id)
	if err != nil {
		return nil, errors.New("User not found")
	}

	return user, nil
}

func (s *userService) UpdateProfile(
	id int,
	name *string,
	image *string,
) error {
	return s.userRepo.UpdateProfile(id, name, image)
}
