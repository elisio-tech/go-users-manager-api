package database

import (
	"go-users-manager-api/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	// = inicializacao != := criacao e inicializacao
	DB, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		return err // manter o erro original facilita a debug
	}
	// cria as tabelas com base no meu struct users
	return DB.AutoMigrate(&models.User{})
}
