package main

import (
	"log"
	"net/http"
	"github.com/gorilla/handlers"
	"./album"
)

func main() {
	router := album.NewRouter() // create routes

	// we need to allow access to our api from the front end
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Content-Type", "Accept-Encoding"})

	// launch server with cors validations
	log.Fatal(http.ListenAndServe(":8000",
	handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router),
	))

}
