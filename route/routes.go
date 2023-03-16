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
	router.GET("/api/set/:id", set.Find)
	router.POST("/api/set", set.Create)
	router.PUT("/api/set/:id", set.Update)
	router.DELETE("/api/set/:id", set.Delete)

	return router
}
