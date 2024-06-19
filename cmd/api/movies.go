package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/ssr0016/todo/internal/data"
)

func (app *application) createMovieFormHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the template file
	tpl := template.Must(template.ParseFiles("./ui/html/pages/create_movie.html"))

	// Execute the template
	if err := tpl.Execute(w, nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")

}

// func (app *application) showMovieFormHandler(w http.ResponseWriter, r *http.Request) {
// 	// Parse the template file
// 	tpl := template.Must(template.ParseFiles("./ui/html/pages/show_detail_movie.html"))

// 	// Execute the template
// 	if err := tpl.Execute(w, nil); err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// }

type Movie struct {
	ID    int64
	Title string
	// Add more fields as needed.
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	// Implement Template parsing
	tpl, err := template.ParseFiles("./ui/html/pages/show_detail_movie.html")
	if err != nil {
		http.Error(w, "Template parsing error", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, movie)
	if err != nil {
		http.Error(w, "Template execution error", http.StatusInternalServerError)
	}

	// // Encode the struct to JSON and send it as the HTTP response.
	// err = app.writeJSON(w, http.StatusOK, movie, nil)
	// if err != nil {
	// 	app.logger.Println(err)
	// 	http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	// }
}
