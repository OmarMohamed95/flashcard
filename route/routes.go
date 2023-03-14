package route

import (
	"flashcards-api/controller/set"

	"github.com/julienschmidt/httprouter"
)

func InitRouter() *httprouter.Router {
	router := httprouter.New()
	router.GET("/set", set.FindAll)
	// router.POST("/set", set.Create)

	return router
}
