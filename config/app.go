package config

import (
	"github.com/alexedwards/scs/v2"
)

// AppConfig is the struct containing main configurations of the application
type AppConfig struct {
	Session *scs.SessionManager
}
