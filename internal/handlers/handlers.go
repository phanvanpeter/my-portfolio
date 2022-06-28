package handlers

import (
	"github.com/phanvanpeter/my-portfolio/internal/render"
	"github.com/phanvanpeter/my-portfolio/models"
	"net/http"
)

// Home renders a Home page
func Home(w http.ResponseWriter, r *http.Request) {
	file := "home.page"

	strMap := make(map[string]string)
	strMap["author"] = "Peter Phanvan"
	render.Template(w, file, &models.TemplateData{
		StringMap: strMap,
	})

}

// About renders an About page
func About(w http.ResponseWriter, r *http.Request) {
	file := "about.page"
	render.Template(w, file, nil)
}
