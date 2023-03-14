package model

import (
	"flashcards-api/database"

	"gorm.io/gorm"
)

type Set struct {
	gorm.Model
	Name string `gorm:"" json:"name"`
}

func init() {
	db := database.Connect()
	db.AutoMigrate(&Set{})
}
