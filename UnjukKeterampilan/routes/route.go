package routes

import (
	"UnjukKeterampilan/controllers"

	echojwt "github.com/labstack/echo-jwt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo) {
	e.Use(middleware.Logger())
	eAuth := e.Group("")
	eAuth.Use(echojwt.JWT([]byte("123")))
	// GET & POST People
	eAuth.POST("/people", controllers.CreatePeopleController)
	eAuth.GET("/people", controllers.GetPeopleController)
	// GET & POST Planet
	eAuth.POST("/planet", controllers.CreatePlanetController)
	eAuth.GET("/planet", controllers.GetPlanetController)
	// GET & POST Starships
	eAuth.POST("/starship", controllers.CreateStarshipController)
	eAuth.GET("/starship", controllers.GetStarshipController)
	//e.POST("/login", controllers.RegisterController)
	e.POST("/register", controllers.RegisterController)
}
