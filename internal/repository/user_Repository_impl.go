package repository

import (
	"go-users-manager-api/internal/models"

	"gorm.io/gorm"
)

type SQLiteUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *SQLiteUserRepository {
	return &SQLiteUserRepository{db: db}
}

func (u *SQLiteUserRepository) Create(usr *models.User) error {
	return u.db.Create(usr).Error
}

func (u *SQLiteUserRepository) List() ([]models.User, error) {
	var usr []models.User

	if err := u.db.Find(&usr).Error; err != nil {
		return nil, err
	}
	return usr, nil
}

func (u *SQLiteUserRepository) FindByID(id string) (*models.User, error) {
	var usr models.User
	if err := u.db.First(&usr, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &usr, nil
}

func (u *SQLiteUserRepository) Update(id string, user models.User) error {
	return u.db.Model(&models.User{}).Where("id = ?", id).Updates(user).Error
}

func (u *SQLiteUserRepository) Delete(id string) error {
	return u.db.Delete(&models.User{}, "id = ?", id).Error
}
