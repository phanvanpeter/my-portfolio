package main

import "net/http"

// SessionLoad loads and saves data for the current request
func SessionLoad(next http.Handler) http.Handler {
	return appConfig.Session.LoadAndSave(next)
}
