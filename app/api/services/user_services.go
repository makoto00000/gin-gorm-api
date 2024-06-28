package services

import (
	"app/api/models"
	"app/api/repositories"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(database *gorm.DB) {
	db = database
}

func GetAllUsers() ([]models.User, error) {
	return repositories.GetAllUsers(db)
}

func CreateUser(user *models.User) error {
	return repositories.CreateUser(db, user)
}