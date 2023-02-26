package config

import (
	"html/template"
	"log"
)

//AppConfig holds the application config , template cache currently
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
}
