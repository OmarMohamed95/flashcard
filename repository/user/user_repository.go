package user

import (
	"flashcards-api/app/database"
	"flashcards-api/model"
	"fmt"
)

type User struct {
	model.User
}

func (s *User) Create() {
	database.DB.Create(&s)
}

func FindBy(whereConditions map[string]string) model.User {
	var user model.User

	for key, val := range whereConditions {
		database.DB = database.DB.Where(fmt.Sprintf("%s = ?", key), val)
	}

	database.DB.Find(&user)

	return user
}
