package api

import (
	"encoding/json"
	"net/http"
)

type JsonResponse struct {
	w *http.ResponseWriter
}

type DataRes struct {
	Data          interface{}
	StatusCode    int16
	StatusMessage string
}

type ErrorRes struct {
	Error      interface{}
	StatusCode int16
}

func Json(w http.ResponseWriter) JsonResponse {
	wrr := JsonResponse{w: &w}
	(*wrr.w).Header().Set("Content-Type", "application/json")

	return wrr
}

func (wrr JsonResponse) Respond(res DataRes) {
	data, err := json.Marshal(res)
	if err != nil {
		wrr.RespondError(ErrorRes{Error: "Error while encoding json", StatusCode: http.StatusInternalServerError})
		return
	}
	(*wrr.w).WriteHeader(int(res.StatusCode))
	(*wrr.w).Write(data)
}

func (wrr JsonResponse) RespondError(errRes ErrorRes) {
	data, err := json.Marshal(errRes)
	if err != nil {
		wrr.RespondError(ErrorRes{Error: "Error while encoding json", StatusCode: http.StatusInternalServerError})
		return
	}
	(*wrr.w).WriteHeader(int(errRes.StatusCode))
	(*wrr.w).Write(data)
}
