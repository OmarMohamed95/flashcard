package card

import (
	"encoding/json"
	"flashcards-api/app/api"
	cardRepository "flashcards-api/repository/card"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var card cardRepository.Card
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		api.Json(w).RespondError(api.ErrorRes{Error: "Error while parsing body", StatusCode: http.StatusBadRequest})
		return
	}

	json.Unmarshal(body, &card)
	card.Create()

	api.Json(w).Respond(api.DataRes{
		Data:          card,
		StatusCode:    http.StatusCreated,
		StatusMessage: "Card created successfully",
	})
}

func FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cards := cardRepository.GetAll()
	api.Json(w).Respond(api.DataRes{
		Data:          cards,
		StatusCode:    http.StatusOK,
		StatusMessage: "Cards retrieved successfully",
	})
}

func Find(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		api.Json(w).RespondError(api.ErrorRes{Error: "Invalid card id", StatusCode: http.StatusBadRequest})
		return
	}

	card := cardRepository.FindById(int64(id))
	if card.ID == 0 {
		api.Json(w).RespondError(api.ErrorRes{Error: "Card not found", StatusCode: http.StatusNotFound})
		return
	}

	api.Json(w).Respond(api.DataRes{
		Data:          card,
		StatusCode:    http.StatusOK,
		StatusMessage: "Card retrieved successfully",
	})
}

func Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		api.Json(w).RespondError(api.ErrorRes{Error: "Invalid card id", StatusCode: http.StatusBadRequest})
		return
	}

	var cardUpdates map[string]interface{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		api.Json(w).RespondError(api.ErrorRes{Error: "Error while parsing body", StatusCode: http.StatusBadRequest})
		return
	}

	json.Unmarshal(body, &cardUpdates)
	card := cardRepository.Update(int64(id), cardUpdates)
	if card.ID == 0 {
		api.Json(w).RespondError(api.ErrorRes{Error: "Card not found", StatusCode: http.StatusNotFound})
		return
	}

	api.Json(w).Respond(api.DataRes{
		Data:          card,
		StatusCode:    http.StatusOK,
		StatusMessage: "Card updated successfully",
	})
}

func Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		api.Json(w).RespondError(api.ErrorRes{Error: "Invalid card id", StatusCode: http.StatusBadRequest})
		return
	}

	cardRepository.Delete(int64(id))
	api.Json(w).Respond(api.DataRes{
		Data:          []int8{},
		StatusCode:    http.StatusOK,
		StatusMessage: "Card deleted successfully",
	})
}
