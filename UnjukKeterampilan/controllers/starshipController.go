package controllers

import (
	"UnjukKeterampilan/models"
	"UnjukKeterampilan/repositories"
	"net/http"

	"UnjukKeterampilan/configs.DB"

	"github.com/labstack/echo/v4"
)

func CreateStarshipController(c echo.Context) error {
	var starshipRequest models.Starship
	c.Bind(&starshipRequest)

	err := repositories.AddStarship(&starshipRequest)
	//result := configs.DB.Create(&newsRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Baseresponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
	}
	return c.JSON(200, starshipRequest)
}

func GetStarshipController(c echo.Context) error {
	var starship []models.Starship
	result := configs.DB.Find(&starship)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.Baseresponse{
			Message: result.Error.Error(),
			Status:  false,
			Data:    nil,
		})
	} // result.Context sebelumnya...
	return c.JSON(http.StatusOK, models.Baseresponse{
		Message: "Berhasil menampilkan data",
		Status:  true,
		Data:    starship,
	}) //awalnya "news" saja
}
