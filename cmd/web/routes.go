package main

import (
	"github.com/go-chi/chi"
	"goAnsible/pkg/config"
	"goAnsible/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
