package route

import (
	"flashcards-api/controller/set"

	"github.com/julienschmidt/httprouter"
)

var router *httprouter.Router

func init() {
	router = httprouter.New()
}

func RegisterRoutes() *httprouter.Router {
	router.GET("/api/set", set.FindAll)
	// router.POST("/set", set.Create)

	return router
}
