package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// the function routes accepts app as a pointer receiver, this allows app to utilize the function routes, similar to methods in OOP
func (app *application) routes() http.Handler {
	// create variable router that uses the httprouter package from github
	router := httprouter.New()

	
	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodGet, "/v1/movie/:id", app.getOneMovie)
	router.HandlerFunc(http.MethodGet, "/v1/movies", app.getAllMovies)

	return app.enableCORS(router)
}
