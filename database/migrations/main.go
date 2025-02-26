package main

import (
	"Gin-GORM-Project/database"
	"Gin-GORM-Project/models"
	"log"
)

func init() {
	database.Connection()
}

func main() {
	err := database.DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("Error auto migrate: ", err)
	}
}
