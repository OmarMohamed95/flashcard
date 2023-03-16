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

func FindById(id int64) model.Set {
	var set model.Set
	db.Where("ID = ?", id).Find(&set)

	return set
}

func Update(id int64, setUpdates map[string]interface{}) model.Set {
	set := FindById(id)
	db.Model(&set).Where("ID = ?", id).Updates(setUpdates)

	return set
}

func Delete(id int64) {
	var set model.Set
	db.Where("ID = ?", id).Delete(&set)
}
