package helpers

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv() {
	errLoadDotenv := godotenv.Load()
	if errLoadDotenv != nil {
		log.Fatalf("Error loading .env file: %s", errLoadDotenv)
	}
}
