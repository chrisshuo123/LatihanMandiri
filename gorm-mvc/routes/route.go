package routes

import (
	"gorm/controllers"

	echojwt "github.com/labstack/echo-jwt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo) {
	e.Use(middleware.Logger())
	eAuth := e.Group("")
	eAuth.Use(echojwt.JWT([]byte("123")))
	eAuth.POST("/news", controllers.CreateNewsController)
	eAuth.GET("/news", controllers.GetNewsController)
	//e.POST("/login", controllers.RegisterController)
	e.POST("/register", controllers.RegisterController)
}
