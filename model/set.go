package model

import (
	"flashcards-api/app/database"

	"gorm.io/gorm"
)

type Set struct {
	gorm.Model
	Name string `gorm:"" json:"name"`
}

func init() {
	database.DB.AutoMigrate(&Set{})
}
