package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const hostAddr = ":8080"
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Server running on a port %s\n", hostAddr)
	err := http.ListenAndServe(hostAddr, nil)
	if err != nil {
		log.Fatal("error running a server", err)
	}
}

// Home renders a Home page
func Home(w http.ResponseWriter, r *http.Request) {
	file := "./templates/home.page.gohtml"
	tmpl, err := template.ParseFiles(file)
	if err != nil {
		log.Fatalf("Error parsing a file %s", file)
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
