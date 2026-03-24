package service

import (
	"go-users-manager-api/internal/models"
	"go-users-manager-api/internal/repository"
)

type UserService struct {
	UserRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{UserRepo: repo}
}

func (u *UserService) GetUsers() ([]models.User, error) {
	users, err := u.UserRepo.List()
	if err != nil {
		return nil, err
	}
	return users, nil
}
