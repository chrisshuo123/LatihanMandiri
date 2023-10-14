package repositories

import (
	"gorm/configs.DB"
	"gorm/models"
)

func Register(user *models.User) error {
	result := configs.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
