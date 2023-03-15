package set

import (
	"flashcards-api/app/database"
	"flashcards-api/model"

	"gorm.io/gorm"
)

var db *gorm.DB

type Set struct {
	model.Set
}

func init() {
	db = database.Connect()
}

func GetAll() []model.Set {
	var sets []model.Set
	db.Find(&sets)

	return sets
}

func (s *Set) Create() {
	db.Create(&s)
}
