package main

import (
	"encoding/json"
	"net/http"
)

/* 	
	the function writeJSON accepts a receiver pointer to app of type application so that app can access the function within the main package
	writeJSON accepts the http response writer, the status code of the response, the data which is 'defaulted' using interface{}, and a wrap
	string to make the JSON output look good.
*/
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, wrap string) error {
	// create wrapper variable and put a map of strings to interfaces inside of it
	wrapper := make(map[string]interface{})

	// put data into the wrapper by putting the wrap parameter as a key and the data as a value
	wrapper[wrap] = data

	// json.Marshal returns the JSON encoding of the wrapper, along with the error. After that, we perform error handling
	js, err := json.Marshal(wrapper)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status) 
	w.Write(js)

	return nil
}

/* 	
	the function errorJSON takes in the response writer as a parameter and the error message
	errorJSON then calls the writeJSON function declared above and passes in the error message so that it is outputted to the browser
*/
func (app *application) errorJSON(w http.ResponseWriter, err error) {
	type jsonError struct {
		Message string `json:"message"`
	}

	theError := jsonError{
		Message: err.Error(),
	}

	app.writeJSON(w, http.StatusBadRequest, theError, "error")
}