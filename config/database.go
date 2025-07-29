package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDatabase() *gorm.DB {
	dsn := "root:abdillah24@tcp(127.0.0.1:3306)/auth_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to Connected server", err)
	}
	db.AutoMigrate()
	return db
}
