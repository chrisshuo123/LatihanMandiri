package configs

import (
	"gorm/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:root@tcp(127.0.0.1:3306)/swapi_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	InitMigration()
}

func InitMigration() {
	// Migrate the Schema
	DB.AutoMigrate(&models.People{}, &models.People{})
	DB.AutoMigrate(&models.Planet{}, &models.Planet{})
	DB.AutoMigrate(&models.Starships{}, &models.Starships{})
}
