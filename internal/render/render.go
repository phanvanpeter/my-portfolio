package render

import (
	"html/template"
	"log"
	"net/http"
)

var TmplFolder = "./templates/"
const tmplExtension = ".gohtml"

// Template renders and sends to browser an HTML from the given template file
func Template(w http.ResponseWriter, file string) {
	tmpl, err := template.ParseFiles(TmplFolder + file + tmplExtension)
	if err != nil {
		log.Fatalf("Error parsing a file %s; %s", file, err)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal("error rendering a template", err)
	}
}
