package set

import (
	"flashcards-api/app/database"
	"flashcards-api/model"
)

type Set struct {
	model.Set
}

func GetAll() []model.Set {
	var sets []model.Set
	database.DB.Find(&sets)

	return sets
}

func (s *Set) Create() {
	database.DB.Create(&s)
}

func FindById(id int64) model.Set {
	var set model.Set
	database.DB.Where("ID = ?", id).Find(&set)

	return set
}

func Update(id int64, setUpdates map[string]interface{}) model.Set {
	set := FindById(id)
	database.DB.Model(&set).Where("ID = ?", id).Updates(setUpdates)

	return set
}

func Delete(id int64) {
	var set model.Set
	database.DB.Where("ID = ?", id).Delete(&set)
}
