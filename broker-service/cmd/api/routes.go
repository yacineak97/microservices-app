package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Application) routes() http.Handler {
	router := chi.NewRouter()

	// Configure CORS settings
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Middleware for health check endpoint
	router.Use(middleware.Heartbeat("/ping"))

	// Define routes
	router.Post("/", app.Broker)

	router.Post("/log-grpc", app.LogViaGRPC)

	router.Post("/handle", app.HandleSubmission)

	return router
}
