package main

import (
	"fmt"
	"log"
	"net/http"
)

const hostAddr = ":8080"
func main() {
	err := run()
	if err != nil {
		log.Fatal("error running a server", err)
	}
}

// run starts the server and runs the web application
func run() error {
	router := Route()

	fmt.Printf("Server running on a port %s\n", hostAddr)
	err := http.ListenAndServe(hostAddr, router)
	if err != nil {
		return err
	}
	return nil
}