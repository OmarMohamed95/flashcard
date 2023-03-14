package main

import (
	"flashcards-api/route"
	"log"
	"net/http"
)

func main() {
	router := route.InitRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
