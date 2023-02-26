package main

import (
	"fmt"
	"goAnsible/pkg/config"
	"goAnsible/pkg/handlers"
	"goAnsible/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8089"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
