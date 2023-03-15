package main

import (
	"flashcards-api/route"
	"log"
	"net/http"
)

func main() {
	router := route.RegisterRoutes()
	log.Fatal(http.ListenAndServe(":8080", router))
}
