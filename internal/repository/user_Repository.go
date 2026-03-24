package repository

import "go-users-manager-api/internal/models"

// define o contrato quais metodos o user pode ter
type UserRepository interface {
	FindAll() (*models.User, error)
	Create(user *models.User) error
	FindByID(id string) error
	GetByEmail(email string) (*models.User, error)
	Update(user *models.User) error
	Delete(id string) error
}
