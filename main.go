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

	router.Use(interceptor.LoggingMiddleware)
	router.Use(interceptor.AuthMiddleware)

	router.HandleFunc("/index", api.Index).Methods("GET")

	router.HandleFunc("/addCustomer", api.AddCustomer).Methods("POST")
	router.HandleFunc("/getAllCustomers", api.GetAllCustomers).Methods("GET")
	router.HandleFunc("/getCustomer/{id}", api.GetCustomerById).Methods("GET")

	router.HandleFunc("/updateCustomer", api.UpdateCustomerData).Methods("PUT")
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	log.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", c.Handler(router)))
}
