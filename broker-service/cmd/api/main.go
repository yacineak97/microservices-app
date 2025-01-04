package main

import (
	"fmt"
	"log"
	"net/http"
)

const serverPort = "80"

type Application struct{}

func main() {
	app := Application{}

	log.Printf("Broker service is now running on port %s\n", serverPort)

	// Create the HTTP server configuration
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", serverPort),
		Handler: app.routes(),
	}

	// Start the server and handle potential errors
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server encountered an error: %v", err)
	}
}
