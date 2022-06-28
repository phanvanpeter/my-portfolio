package render

import (
	"github.com/phanvanpeter/my-portfolio/models"
	"html/template"
	"log"
	"net/http"
)

var TmplFolder = "./templates/"
const tmplExtension = ".gohtml"

// Template renders and sends to browser an HTML from the given template file
func Template(w http.ResponseWriter, file string, td *models.TemplateData) {
	tmpl, err := template.ParseFiles(TmplFolder + file + tmplExtension)
	if err != nil {
		log.Fatalf("Error parsing a file %s; %s", file, err)
	}

	err = tmpl.Execute(w, td)
	if err != nil {
		log.Fatal("error rendering a template", err)
	}
}
