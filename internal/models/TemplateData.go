package models

import "github.com/mihailtudos/bookings/internal/forms"

// TemplateData holds data sent from handlers to the template
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[int]int
	FlotMap   map[float32]float32
	Data      map[string]any
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}
