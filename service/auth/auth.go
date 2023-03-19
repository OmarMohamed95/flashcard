package auth

import (
	"flashcards-api/app/config"
	"flashcards-api/model"

	"github.com/golang-jwt/jwt"
)

func NewToken(user model.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      user.ID,
		"email":    user.Email,
		"username": user.Username,
	})

	signedToken, _ := token.SignedString([]byte(config.Get("SECRET")))

	return signedToken
}
