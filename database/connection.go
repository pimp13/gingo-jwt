package database

import (
	"Gin-GORM-Project/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connection() {
	var errConnection error
	var dns string

	if config.DbDriver == "mysql" {
		dns = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.DbUser,
			config.DbPassword,
			config.DbHost,
			config.DbPort,
			config.DbName,
		)
		DB, errConnection = gorm.Open(mysql.Open(dns), &gorm.Config{})
	}
	if config.DbDriver == "pgsql" {
		dns = fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
			config.DbHost,
			config.DbUser,
			config.DbPassword,
			config.DbName,
			config.DbPort,
		)
		DB, errConnection = gorm.Open(postgres.Open(dns), &gorm.Config{})
	}
	if errConnection != nil {
		panic("Failed to connect to the database!")
	}

	log.Println("Successfully connected to the database!")
}
