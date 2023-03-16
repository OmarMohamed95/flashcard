package route

import (
	setController "flashcards-api/controller/set"

	"github.com/julienschmidt/httprouter"
)

var router *httprouter.Router

func init() {
	router = httprouter.New()
}

func RegisterRoutes() *httprouter.Router {
	router.GET("/api/set", setController.FindAll)
	router.GET("/api/set/:id", setController.Find)
	router.POST("/api/set", setController.Create)
	router.PUT("/api/set/:id", setController.Update)
	router.DELETE("/api/set/:id", setController.Delete)

	return router
}
