package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
	"log"
)

//AppConfig holds the application config , template cache currently
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager // this is where we are storing the session so non main package can access
}
