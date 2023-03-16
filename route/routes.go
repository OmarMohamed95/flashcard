package route

import (
	cardController "flashcards-api/controller/card"
	setController "flashcards-api/controller/set"

	"github.com/julienschmidt/httprouter"
)

var router *httprouter.Router

func init() {
	router = httprouter.New()
}

func RegisterRoutes() *httprouter.Router {
	// Set
	router.GET("/api/set", setController.FindAll)
	router.GET("/api/set/:id", setController.Find)
	router.POST("/api/set", setController.Create)
	router.PUT("/api/set/:id", setController.Update)
	router.DELETE("/api/set/:id", setController.Delete)

	// Card
	router.GET("/api/card", cardController.FindAll)
	router.GET("/api/card/:id", cardController.Find)
	router.POST("/api/card", cardController.Create)
	router.PUT("/api/card/:id", cardController.Update)
	router.DELETE("/api/card/:id", cardController.Delete)

	return router
}
