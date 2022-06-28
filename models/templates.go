package models

type TemplateData struct {
	Data      map[string]interface{}
	StringMap map[string]string
	CSRFToken string
}
