package model

import (
	"flashcards-api/app/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"" json:"username"`
	Email    string `gorm:"" json:"email"`
	Password string `gorm:"" json:"password"`
}

func init() {
	database.DB.AutoMigrate(&User{})
}
