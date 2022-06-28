package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/phanvanpeter/my-portfolio/internal/handlers"
	"log"
	"net/http"
)

const hostAddr = ":8080"
func main() {
	handler := Route()

	fmt.Printf("Server running on a port %s\n", hostAddr)
	err := http.ListenAndServe(hostAddr, handler)
	if err != nil {
		log.Fatal("error running a server", err)
	}
}

func Route() http.Handler {
	h := chi.NewRouter()

	h.Use(middleware.Recoverer)

	h.Get("/", handlers.Home)
	h.Get("/about", handlers.About)

	return h
}
