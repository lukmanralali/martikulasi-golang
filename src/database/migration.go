package database

import (
	"../models"
	"fmt"
)

func Migrate() {
	fmt.Println("Doing AutoMigrate Gorm")
	db := GetConnection()
	db.AutoMigrate(&models.User{}, &models.UrlShortCode{})
}
