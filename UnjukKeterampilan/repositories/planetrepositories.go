package repositories

import (
	"UnjukKeterampilan/configs.DB"
	"UnjukKeterampilan/models"
)

func GetPlanet(planetList *[]models.Planet) error {
	result := configs.DB.Find(&planetList)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AddPlanet(planetDB *models.Planet) error {
	result := configs.DB.Create(&planetDB)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
