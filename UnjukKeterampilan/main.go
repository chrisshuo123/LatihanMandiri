package main

import (
	"UnjukKeterampilan/configs.DB"
	"UnjukKeterampilan/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	//configs.initDatabase()
	configs.InitDatabase()
	e := echo.New()
	//configs.initRoute()
	routes.InitRoute(e)
	e.Start(":8000")
	//e.Start("https://swapi.dev/api")
}

/*type Product struct {
	gorm.Model
	Code  string
	Price uint
} */

/*
func main() {
	// Create a new user
	user := &Product{Code: "D42", Price: 100}
	db.Create(user)
	// Read
	var product Product
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42
	// Update - update product's price to 200
	db.Model(&product).Update("price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // Non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"price": 200, "code": "F42"})
	// Delete
	db.Delete(&product, 1)
} */

// connect gorm mysql
