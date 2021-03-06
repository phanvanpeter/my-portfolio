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
	r.Use(CSRF)
	r.Use(SessionLoad)

	r.Get("/", handlers.Home)
	r.Get("/about", handlers.About)

	r.Get("/tasks", handlers.Tasks)
	r.Post("/tasks", handlers.PostTask)
	r.Post("/tasks/{id}/complete", handlers.CompleteTask)
	r.Post("/tasks/{id}/edit", handlers.EditTask)
	r.Post("/tasks/{id}/delete", handlers.DeleteTask)

	return r
}
