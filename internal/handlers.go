package internal

import (
	"html/template"
	"log"
	"net/http"
)

var TmplFolder = "./templates/"

// Home renders a Home page
func Home(w http.ResponseWriter, r *http.Request) {
	file := "home.page.gohtml"
	tmpl, err := template.ParseFiles(TmplFolder + file)
	if err != nil {
		log.Fatalf("Error parsing a file %s; %s", file, err)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal("error rendering a template", err)
	}
}

// About renders an About page
func About(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("About page"))
	if err != nil {
		log.Fatal(err)
	}
}
