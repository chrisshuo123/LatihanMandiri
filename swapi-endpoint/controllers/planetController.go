package controllers

import (
	"gorm/configs.DB"
	"gorm/models"
	"gorm/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreatePlanetController(c echo.Context) error {
	var planetRequest models.Planet
	c.Bind(&planetRequest)

	err := repositories.AddPlanet(&planetRequest)
	//result := configs.DB.Create(&newsRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Baseresponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
	}
	return c.JSON(200, planetRequest)
}

func GetPlanetController(c echo.Context) error {
	var planet []models.Planet
	result := configs.DB.Find(&planet)
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
		Data:    planet,
	}) //awalnya "news" saja
}
