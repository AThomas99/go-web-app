package main

import (
	"hello-world/pkg/config"
	"hello-world/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes (app *config.AppConfig) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(NoSurf)

	r.Get("/", handlers.Repo.Home)
	r.Get("/about", handlers.Repo.About)

	return r
}