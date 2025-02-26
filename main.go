package main

import (
	"Gin-GORM-Project/bootstrap"
	"Gin-GORM-Project/database"
	"Gin-GORM-Project/helpers"
)

func init() {
	database.Connection()
	helpers.LoadEnv()
}

func main() {
	bootstrap.App()
}
