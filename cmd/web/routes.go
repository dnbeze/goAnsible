package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"goAnsible/pkg/config"
	"goAnsible/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
