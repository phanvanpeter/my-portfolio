package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/phanvanpeter/my-portfolio/internal/handlers"
	"net/http"
)

// Route creates a router which processes the HTTP requests
func Route() http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(SessionLoad)

	r.Get("/", handlers.Home)
	r.Get("/about", handlers.About)

	return r
}
