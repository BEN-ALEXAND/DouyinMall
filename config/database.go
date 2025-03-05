package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"dymall/server/models"
)

var DB *gorm.DB

func InitDB() {
	dsn := "root:xm0322@tcp(127.0.0.1:3306)/dymall?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(
		&models.User{}, 
		&models.Product{}, 
		&models.Order{},
		&models.Cart{},
	)
}