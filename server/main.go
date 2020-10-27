package main

import (
	"log"
	"net/http"

	"github.com/bygui86/go-graphql/server/book"

	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()
	router = book.RegisterRoutes(router)

	log.Println("Server ready at 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
