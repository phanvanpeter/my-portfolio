package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/phanvanpeter/my-portfolio/config"
	"log"
	"net/http"
	"time"
)

const hostAddr = ":8080"

var appConfig *config.AppConfig
var session *scs.SessionManager

func main() {
	err := run()
	if err != nil {
		log.Fatal("error running a server", err)
	}
}

// run starts the server and runs the web application
func run() error {
	session = initSession()

	appConfig = &config.AppConfig{
		Session: session,
	}

	router := Route()

	fmt.Printf("Server running on a port %s\n", hostAddr)
	err := http.ListenAndServe(hostAddr, router)
	if err != nil {
		return err
	}
	return nil
}

func initSession() *scs.SessionManager {
	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Persist = true
	session.Cookie.Secure = true

	return session

}
