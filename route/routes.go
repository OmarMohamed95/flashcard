package route

import (
	"flashcards-api/app/security"
	authController "flashcards-api/controller/auth"
	cardController "flashcards-api/controller/card"
	setController "flashcards-api/controller/set"

	"github.com/julienschmidt/httprouter"
)

var router *httprouter.Router

func init() {
	router = httprouter.New()
}

func RegisterRoutes() *httprouter.Router {
	// Auth
	router.POST("/api/login", authController.Login)
	router.POST("/api/register", authController.Register)

	// Set
	router.GET("/api/set", security.Auth(setController.FindAll))
	router.GET("/api/set/:id", security.Auth(setController.Find))
	router.POST("/api/set", security.Auth(setController.Create))
	router.PUT("/api/set/:id", security.Auth(setController.Update))
	router.DELETE("/api/set/:id", security.Auth(setController.Delete))

	// Card
	router.GET("/api/card", security.Auth(cardController.FindAll))
	router.GET("/api/card/:id", security.Auth(cardController.Find))
	router.POST("/api/card", security.Auth(cardController.Create))
	router.PUT("/api/card/:id", security.Auth(cardController.Update))
	router.DELETE("/api/card/:id", security.Auth(cardController.Delete))

	return router
}
