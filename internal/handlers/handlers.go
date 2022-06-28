package handlers

import (
	"github.com/phanvanpeter/my-portfolio/config"
	"github.com/phanvanpeter/my-portfolio/internal/render"
	"github.com/phanvanpeter/my-portfolio/internal/repository"
	"net/http"
)

var (
	app *config.AppConfig
	db  repository.DBRepository
)

func InitHandlers(a *config.AppConfig, dbRepo repository.DBRepository) {
	app = a
	db = dbRepo
}

var name = ""

// Home renders a Home page
func Home(w http.ResponseWriter, r *http.Request) {
	file := "home.page"
	author := "Peter Phanvan"

	strMap := map[string]string{}
	strMap["author"] = author

	app.Session.Put(r.Context(), "author", author)

	render.Template(w, r, file, &config.TemplateData{
		StringMap: strMap,
	})
}

// About renders an About page
func About(w http.ResponseWriter, r *http.Request) {
	file := "about.page"

	author := app.Session.GetString(r.Context(), "author")

	strMap := map[string]string{
		"author": author,
	}

	render.Template(w, r, file, &config.TemplateData{
		StringMap: strMap,
	})
	app.Session.Remove(r.Context(), "author")
}
