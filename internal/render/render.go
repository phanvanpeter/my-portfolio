package render

import (
	"github.com/justinas/nosurf"
	"github.com/phanvanpeter/my-portfolio/models"
	"html/template"
	"log"
	"net/http"
)

var TmplFolder = "./templates/"
const tmplExtension = ".gohtml"

func AddDefaultData(r *http.Request, td *models.TemplateData) {
	td.CSRFToken = nosurf.Token(r)
}

// Template renders and sends to browser an HTML from the given template file
func Template(w http.ResponseWriter, r *http.Request, file string, td *models.TemplateData) {
	tmpl, err := template.ParseFiles(TmplFolder + file + tmplExtension)
	if err != nil {
		log.Fatalf("Error parsing a file %s; %s", file, err)
	}

	AddDefaultData(r, td)

	err = tmpl.Execute(w, td)
	if err != nil {
		log.Fatal("error rendering a template", err)
	}
}
