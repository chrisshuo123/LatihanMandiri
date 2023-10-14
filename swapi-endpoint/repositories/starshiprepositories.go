package repositories

import (
	"gorm/configs.DB"
	"gorm/models"
)

func GetStarship(starshipList *[]models.Starships) error {
	result := configs.DB.Find(&starshipList)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AddStarship(starshipDB *models.Starships) error {
	result := configs.DB.Create(&starshipDB)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
