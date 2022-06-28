package config

import "github.com/phanvanpeter/my-portfolio/internal/forms"

type TemplateData struct {
	Data       map[string]interface{}
	StringMap  map[string]string
	FormErrors forms.FormErrors
	CSRFToken  string
}
