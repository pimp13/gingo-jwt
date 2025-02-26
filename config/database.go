package config

import (
	"Gin-GORM-Project/helpers"
	"os"
)

var (
	DbDriver   string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
)

func init() {
	helpers.LoadEnv()

	DbDriver = os.Getenv("DB_DRIVER")
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbUser = os.Getenv("DB_USER")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbName = os.Getenv("DB_NAME")
}
