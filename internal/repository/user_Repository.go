package repository

import "go-users-manager-api/internal/models"

// define o contrato quais metodos o user pode ter
type UserRepository interface {
	List() ([]models.User, error)
	Create(user *models.User) error
	FindByID(id string) (*models.User, error)
	Update(id string, usr *models.User) error
	Delete(id string) error
}
