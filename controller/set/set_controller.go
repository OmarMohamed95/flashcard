package set

import (
	"encoding/json"
	"flashcards-api/app/api"
	"flashcards-api/repository/set"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var set set.Set
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		api.Json(w).RespondError(api.ErrorRes{Error: "Error while parsing body", StatusCode: http.StatusBadRequest})
	}

	json.Unmarshal(body, &set)
	set.Create()

	api.Json(w).Respond(api.DataRes{
		Data:          set,
		StatusCode:    http.StatusCreated,
		StatusMessage: "Set created successfully",
	})
}

func FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sets := set.GetAll()
	api.Json(w).Respond(api.DataRes{
		Data:          sets,
		StatusCode:    http.StatusOK,
		StatusMessage: "Data retrieved successfully",
	})
}
