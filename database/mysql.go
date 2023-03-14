package database

import (
	"flashcards-api/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	user := config.Get("DB_USER")
	password := config.Get("DB_PASSWORD")
	host := config.Get("DB_HOST")
	port := config.Get("DB_PORT")
	dbName := config.Get("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to DB")
	}

	return db
}
