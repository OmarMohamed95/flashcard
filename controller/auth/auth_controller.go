package auth

import (
	"encoding/json"
	"flashcards-api/app/api"
	userRepository "flashcards-api/repository/user"
	"flashcards-api/service/auth"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
)

type credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user userRepository.User
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		api.Json(w).RespondError(api.ErrorRes{Error: "Error while parsing body", StatusCode: http.StatusBadRequest})
		return
	}

	json.Unmarshal(body, &user)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		api.Json(w).RespondError(api.ErrorRes{Error: "Internal server error", StatusCode: http.StatusInternalServerError})
	}
	user.Password = string(hashedPassword)

	signedToken := auth.NewToken(user.User)

	user.Create()
	user.Password = ""

	api.Json(w).Respond(api.DataRes{
		Data:          map[string]interface{}{"user": user, "token": signedToken},
		StatusCode:    http.StatusCreated,
		StatusMessage: "User created successfully",
	})
}

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var credentials credentials
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		api.Json(w).RespondError(api.ErrorRes{Error: "Error while parsing body", StatusCode: http.StatusBadRequest})
		return
	}

	json.Unmarshal(body, &credentials)
	user := userRepository.FindBy(map[string]string{"email": credentials.Email})
	if user.ID == 0 {
		api.Json(w).RespondError(api.ErrorRes{Error: "Wrong credentials", StatusCode: http.StatusBadRequest})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		api.Json(w).RespondError(api.ErrorRes{Error: "Wrong credentials", StatusCode: http.StatusInternalServerError})
		return
	}

	signedToken := auth.NewToken(user)

	api.Json(w).Respond(api.DataRes{
		Data:          map[string]string{"token": signedToken},
		StatusCode:    http.StatusOK,
		StatusMessage: "Token retrieved successfully",
	})
}
