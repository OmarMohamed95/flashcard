package route

import (
	"flashcards-api/app/security"
	authController "flashcards-api/controller/auth"
	cardController "flashcards-api/controller/card"
	setController "flashcards-api/controller/set"

	"github.com/go-chi/chi/v5"
)

var router *chi.Mux

func init() {
	router = chi.NewRouter()
}

func RegisterRoutes() *chi.Mux {
	handleAPI()

	return router
}

func handleAPI() {
	// Auth
	router.Post("/api/login", authController.Login)
	router.Post("/api/register", authController.Register)

	// Secured Routes
	router.Route("/api", func(router chi.Router) {
		router.Use(security.Auth)

		// Set
		router.Get("/set", setController.FindAll)
		router.Get("/set/{id}", setController.Find)
		router.Post("/set", setController.Create)
		router.Put("/set/{id}", setController.Update)
		router.Delete("/set/{id}", setController.Delete)

		// Card
		router.Get("/card", cardController.FindAll)
		router.Get("/card/{id}", cardController.Find)
		router.Post("/card", cardController.Create)
		router.Put("/card/{id}", cardController.Update)
		router.Delete("/card/{id}", cardController.Delete)
	})
}
