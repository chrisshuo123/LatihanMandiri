package controllers

import (
	"net/http"

	"UnjukKeterampilan/configs.DB"
	"UnjukKeterampilan/models"
	"UnjukKeterampilan/repositories"

	"github.com/labstack/echo/v4"
)

func CreatePeopleController(c echo.Context) error {
	var peopleRequest models.People
	c.Bind(&peopleRequest)

	err := repositories.AddPeople(&peopleRequest)
	//result := configs.DB.Create(&newsRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Baseresponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
	}
	return c.JSON(200, peopleRequest)
}

func GetPeopleController(c echo.Context) error {
	var news []models.People
	result := configs.DB.Find(&news)
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
		Data:    news,
	}) //awalnya "news" saja
}
