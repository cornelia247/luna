package models

type TemplateData struct {
	StringMap map[string]string
	IntMAp    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}