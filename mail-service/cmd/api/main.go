package main

import (
	"fmt"
	"log"
	"net/http"
)

type Application struct {
}

const webPort = "80"

func main() {
	app := Application{}

	log.Println("Starting mail service on port", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
