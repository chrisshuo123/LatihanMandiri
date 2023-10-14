package controllers

import (
	"gorm/configs.DB"
	"gorm/models"
	"gorm/repositories"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateNewsController(c echo.Context) error {
	var newsRequest models.News
	c.Bind(&newsRequest)

	err := repositories.AddNews(&newsRequest)
	//result := configs.DB.Create(&newsRequest)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Baseresponse{
			Message: err.Error(),
			Status:  false,
			Data:    nil,
		})
	}
	return c.JSON(200, newsRequest)
}

func GetNewsController(c echo.Context) error {
	var news []models.News
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
