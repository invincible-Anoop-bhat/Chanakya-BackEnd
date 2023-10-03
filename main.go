package main

import (
	"Chanakya-BackEnd/api"
	"Chanakya-BackEnd/interceptor"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/index", api.Index).Methods("GET")
	router.HandleFunc("/getAllCustomers", api.GetAllCustomers).Methods("GET")
	router.HandleFunc("/getCustomer/{id}", api.GetCustomerById).Methods("GET")
	router.HandleFunc("/addCustomer", api.AddCustomer).Methods("POST")
	router.Use(interceptor.LoggingMiddleware)
	router.Use(interceptor.AuthMiddleware)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	log.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", c.Handler(router)))
	log.Printf("Server Started on port 8080")
}
