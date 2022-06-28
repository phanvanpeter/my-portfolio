package handlers

import (
	"github.com/phanvanpeter/my-portfolio/internal/render"
	"net/http"
)

// Home renders a Home page
func Home(w http.ResponseWriter, r *http.Request) {
	file := "home.page"
	render.Template(w, file)

}

// About renders an About page
func About(w http.ResponseWriter, r *http.Request) {
	file := "about.page"
	render.Template(w, file)
}
