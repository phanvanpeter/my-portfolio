package main

import (
	"github.com/go-chi/chi/v5"
	"testing"
)

func TestRoute(t *testing.T) {
	router := Route()

	switch r := router.(type) {
	case *chi.Mux:

	default:
		t.Errorf("Expected type *chi.Mux, got %T", r)
	}
}
