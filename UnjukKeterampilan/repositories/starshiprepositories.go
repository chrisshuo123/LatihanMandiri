package repositories

import (
	"UnjukKeterampilan/configs.DB"
	"UnjukKeterampilan/models"
)

func GetStarship(starshipList *[]models.Starship) error {
	result := configs.DB.Find(&starshipList)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AddStarship(starshipDB *models.Starship) error {
	result := configs.DB.Create(&starshipDB)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
