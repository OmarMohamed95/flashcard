package set

import (
	"flashcards-api/app/api"
	"flashcards-api/repository/set"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sets := set.GetAll()
	api.Json(w).Respond(api.DataRes{
		Data:          sets,
		StatusCode:    http.StatusOK,
		StatusMessage: "Data retrieved successfully",
	})
}
