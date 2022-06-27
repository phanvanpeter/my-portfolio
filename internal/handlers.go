package internal

import (
	"html/template"
	"log"
	"net/http"
)

var TmplFolder = "./templates/"
const tmplExtenstion = ".gohtml"

// Home renders a Home page
func Home(w http.ResponseWriter, r *http.Request) {
	file := "home.page"
	RenderTemplate(w, file)

}

// About renders an About page
func About(w http.ResponseWriter, r *http.Request) {
	file := "about.page"
	RenderTemplate(w, file)
}

// RenderTemplate renders and sends to browser an HTML from the given template file
func RenderTemplate(w http.ResponseWriter, file string) {
	tmpl, err := template.ParseFiles(TmplFolder + file + tmplExtenstion)
	if err != nil {
		log.Fatalf("Error parsing a file %s; %s", file, err)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal("error rendering a template", err)
	}
}


