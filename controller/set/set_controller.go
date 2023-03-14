package set

import (
	"encoding/json"
	"flashcards-api/repository/set"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	sets := set.GetAll()
	err := json.NewEncoder(w).Encode(sets)
	if err != nil {
		res, _ := json.Marshal("Internal server error")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(res)
	}
}

func FindAll(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	sets := set.GetAll()
	err := json.NewEncoder(w).Encode(sets)
	if err != nil {
		res, _ := json.Marshal("Internal server error")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(res)
	}
}
