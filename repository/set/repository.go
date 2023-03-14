package set

import (
	"flashcards-api/database"
	"flashcards-api/model"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = database.Connect()
}

func GetAll() []model.Set {
	var sets []model.Set
	db.Find(&sets)

	return sets
}
