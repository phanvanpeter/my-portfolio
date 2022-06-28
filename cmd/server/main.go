package main

import (
	"fmt"
	"github.com/phanvanpeter/my-portfolio/internal/handlers"
	"log"
	"net/http"
)

const hostAddr = ":8080"
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Server running on a port %s\n", hostAddr)
	err := http.ListenAndServe(hostAddr, nil)
	if err != nil {
		log.Fatal("error running a server", err)
	}
}

