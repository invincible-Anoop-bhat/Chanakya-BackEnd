package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	fmt.Println("Welcome to user operations microservice")
	router := mux.NewRouter()

	// Interceptors
	// router.Use(interceptor.LoggingMiddleware)

	// API Handlers
	router.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("Welcome to User Operations Backend!")) }).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	log.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", c.Handler(router)))
}
