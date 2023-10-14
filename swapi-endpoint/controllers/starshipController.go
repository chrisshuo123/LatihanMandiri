package controllers

import (
	"gorm/configs.DB"
	"gorm/models"
	"gorm/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateStarshipController(c echo.Context) error {
	var starshipRequest models.Starships
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
	var starships []models.Starships
	result := configs.DB.Find(&starships)
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
		Data:    starships,
	}) //awalnya "news" saja
}
