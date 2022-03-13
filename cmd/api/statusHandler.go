package main

import (
	"encoding/json"
	"net/http"
)

// these are called receiver functions and they are basically Go's approach to achieving the same goal as OOP
// here's a link to learn more about pointer receivers https://go.dev/tour/methods/4

// statusHandler is a function that writes the current status of the API to the user
func (app *application) statusHandler(w http.ResponseWriter, r *http.Request) {
	currentStatus := AppStatus{
		Status:      "Available",
		Environment: app.config.env,
		Version:     version,
	}

	// MarshalIndent is a function that returns a json and error. This function is used to transform the data type into a JSON.  
	js, err := json.MarshalIndent(currentStatus, "", "\t")
	if err != nil {
		app.logger.Println(err)
	}
	// 'w' is the http response writer which gives utilities to write the response to the browser.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}