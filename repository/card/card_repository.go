package card

import (
	"flashcards-api/app/database"
	"flashcards-api/model"

	"gorm.io/gorm"
)

var db *gorm.DB

type Card struct {
	model.Card
}

func init() {
	db = database.Connect()
}

func GetAll() []model.Card {
	var cards []model.Card
	db.Find(&cards)

	return cards
}

func (s *Card) Create() {
	db.Create(&s)
}

func FindById(id int64) model.Card {
	var card model.Card
	db.Where("ID = ?", id).Find(&card)

	return card
}

func Update(id int64, cardUpdates map[string]interface{}) model.Card {
	card := FindById(id)
	db.Model(&card).Where("ID = ?", id).Updates(cardUpdates)

	return card
}

func Delete(id int64) {
	var card model.Card
	db.Where("ID = ?", id).Delete(&card)
}
