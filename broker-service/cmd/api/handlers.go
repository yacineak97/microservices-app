package main

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (app *Application) Broker(w http.ResponseWriter, r *http.Request) {
	response := jsonResponse{
		Error:   false,
		Message: "Broker endpoint reached successfully",
	}

	output, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		http.Error(w, "Unable to process the request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	_, _ = w.Write(output)
}
