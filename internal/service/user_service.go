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

func (u *UserService) GetUserByID(id string) (*models.User, error) {
	user, err := u.UserRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserService) DeleteUser(id string) error {
	err := u.UserRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) UpdateUser(id string, user *models.User) error {
	return u.UserRepo.Update(id, user)
}
