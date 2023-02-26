package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"goAnsible/pkg/config"
	"goAnsible/pkg/handlers"
	"goAnsible/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8089"

func main() {
	var app config.AppConfig

	session := scs.New()              //scs package handles sessions
	session.Lifetime = 24 * time.Hour // set lifetime of session to expire 24 hr
	session.Cookie.Persist = true     // session will persist a browser being closed - false for fast dying sesion
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false // TODO MUST BE TRUE IN PRODUCTION

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
