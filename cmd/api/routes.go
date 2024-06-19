package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func (app *application) routes() http.Handler {
	// Create a new Chi router.
	r := chi.NewRouter()

	// Middleware setup (optional but recommended).
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	staticPath := "./ui/static"
	fileServer := http.FileServer(http.Dir(staticPath))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// Client Side Rendering
	r.Get("/v1/movies/create", app.createMovieFormHandler)

	// Routes setup.
	r.Get("/v1/healthcheck", app.healthcheckHandler)

	r.Post("/v1/movies", app.createMovieHandler)
	r.Get("/v1/movies/{id}", app.showMovieHandler)

	// Return the Chi router as http.Handler.
	return r
}
