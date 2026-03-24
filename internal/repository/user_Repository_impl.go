package repository

import (
	"go-users-manager-api/internal/models"

	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{db: db}
}

func (u *GormUserRepository) Create(usr *models.User) error {
	return u.db.Create(usr).Error
}

func (u *GormUserRepository) FindAll() ([]models.User, error) {
	var usr []models.User

	if err := u.db.Find(&usr).Error; err != nil {
		return nil, err
	}
	return usr, nil
}
