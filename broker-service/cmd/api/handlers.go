package main

import (
	"net/http"
)

func (app *Application) Broker(w http.ResponseWriter, r *http.Request) {
	response := jsonResponse{
		Error:   false,
		Message: "Broker endpoint reached successfully",
	}

	_ = app.writeJSON(w, http.StatusOK, response)
}
