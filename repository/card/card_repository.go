package card

import (
	"flashcards-api/app/database"
	"flashcards-api/model"
)

type Card struct {
	model.Card
}

func GetAll() []model.Card {
	var cards []model.Card
	database.DB.Find(&cards)

	return cards
}

func (s *Card) Create() {
	database.DB.Create(&s)
}

func FindById(id int64) model.Card {
	var card model.Card
	database.DB.Where("ID = ?", id).Find(&card)

	return card
}

func Update(id int64, cardUpdates map[string]interface{}) model.Card {
	card := FindById(id)
	database.DB.Model(&card).Where("ID = ?", id).Updates(cardUpdates)

	return card
}

func Delete(id int64) {
	var card model.Card
	database.DB.Where("ID = ?", id).Delete(&card)
}
