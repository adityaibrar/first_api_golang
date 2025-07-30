package config

import (
	"first_api_golang/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	dsn := "root:abdillah24@tcp(127.0.0.1:3306)/e_commerce?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to Connected server", err)
	}
	db.AutoMigrate(&models.User{}, &models.Product{})
	return db
}
