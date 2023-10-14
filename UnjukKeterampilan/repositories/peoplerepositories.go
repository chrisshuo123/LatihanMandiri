package repositories

import (
	"UnjukKeterampilan/configs.DB"
	"UnjukKeterampilan/models"
)

func GetPeople(peopleList *[]models.People) error {
	result := configs.DB.Find(&peopleList)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AddPeople(peopleDB *models.People) error {
	result := configs.DB.Create(&peopleDB)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
