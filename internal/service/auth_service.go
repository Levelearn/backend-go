package service

import (
	"errors"
	"fmt"
	"levelearn-backend/internal/entity"
	"levelearn-backend/internal/repository"
	"levelearn-backend/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(username, password string) (string, error)
	Register(user *entity.User) error
}

type authService struct {
	userRepo repository.UserRepository
}

func NewAuthService(userRepo repository.UserRepository) AuthService {
	return &authService{userRepo}
}

func (s *authService) Login(username, password string) (string, error) {
	user, err := s.userRepo.FindByUsername(username)

	fmt.Println("INPUT USERNAME:", username)
	fmt.Println("INPUT PASSWORD:", password)
	fmt.Println("HASH DB:", user.Password)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	return utils.GenerateToken(user.ID, string(user.Role))
}

func (s *authService) Register(user *entity.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashed)
	return s.userRepo.Create(user)
}
