package main

import (
	"github.com/justinas/nosurf"
	"net/http"
)

// SessionLoad loads and saves data for the current request
func SessionLoad(next http.Handler) http.Handler {
	return appConfig.Session.LoadAndSave(next)
}

func CSRF(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		Path: "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure: true,
	})

	return csrfHandler
}
