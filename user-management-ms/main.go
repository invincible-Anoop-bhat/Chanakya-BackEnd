package main

import (
	"fmt"
	"log"
	"net/http"
	"user-management/interceptor"

	"user-management/api"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	fmt.Println("Welcome to User management repo!")

	router := mux.NewRouter()

	// Interceptors
	router.Use(interceptor.LoggingMiddleware)

	// API Handlers
	api.MapHandlersToRoutes(router)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	log.Printf("Starting server at port 8001\n")
	log.Fatal(http.ListenAndServe(":8001", c.Handler(router)))
}
