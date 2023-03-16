package set

import (
	"encoding/json"
	"flashcards-api/app/api"
	setRepository "flashcards-api/repository/set"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var set setRepository.Set
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		api.Json(w).RespondError(api.ErrorRes{Error: "Error while parsing body", StatusCode: http.StatusBadRequest})
		return
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
	sets := setRepository.GetAll()
	api.Json(w).Respond(api.DataRes{
		Data:          sets,
		StatusCode:    http.StatusOK,
		StatusMessage: "Data retrieved successfully",
	})
}

func Find(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		api.Json(w).RespondError(api.ErrorRes{Error: "Invalid set id", StatusCode: http.StatusBadRequest})
		return
	}

	set := setRepository.FindById(int64(id))
	if set.ID == 0 {
		api.Json(w).RespondError(api.ErrorRes{Error: "Set not found", StatusCode: http.StatusNotFound})
		return
	}

	api.Json(w).Respond(api.DataRes{
		Data:          set,
		StatusCode:    http.StatusOK,
		StatusMessage: "Set retrieved successfully",
	})
}

func Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		api.Json(w).RespondError(api.ErrorRes{Error: "Invalid set id", StatusCode: http.StatusBadRequest})
		return
	}

	var setUpdates map[string]interface{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		api.Json(w).RespondError(api.ErrorRes{Error: "Error while parsing body", StatusCode: http.StatusBadRequest})
		return
	}

	json.Unmarshal(body, &setUpdates)
	set := setRepository.Update(int64(id), setUpdates)
	if set.ID == 0 {
		api.Json(w).RespondError(api.ErrorRes{Error: "Set not found", StatusCode: http.StatusNotFound})
		return
	}

	api.Json(w).Respond(api.DataRes{
		Data:          set,
		StatusCode:    http.StatusOK,
		StatusMessage: "Set updated successfully",
	})
}

func Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		api.Json(w).RespondError(api.ErrorRes{Error: "Invalid set id", StatusCode: http.StatusBadRequest})
		return
	}

	setRepository.Delete(int64(id))
	api.Json(w).Respond(api.DataRes{
		Data:          []int8{},
		StatusCode:    http.StatusOK,
		StatusMessage: "Set deleted successfully",
	})
}
