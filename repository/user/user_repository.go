package user

import (
	"flashcards-api/app/database"
	"flashcards-api/model"
	"fmt"

	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	model.User
}

func init() {
	db = database.Connect()
}

func (s *User) Create() {
	db.Create(&s)
}

func FindBy(whereConditions map[string]string) model.User {
	var user model.User

	for key, val := range whereConditions {
		db = db.Where(fmt.Sprintf("%s = ?", key), val)
	}

	db.Find(&user)

	return user
}
