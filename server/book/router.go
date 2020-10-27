package book

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/graphql-go/handler"
)

// TODO try to remove return '*chi.Mux'
func RegisterRoutes(router *chi.Mux) *chi.Mux {

	graphQL := handler.New(&handler.Config{
		Schema:   &Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	router.Use(middleware.Logger)
	router.Handle("/query", graphQL)

	return router
}
