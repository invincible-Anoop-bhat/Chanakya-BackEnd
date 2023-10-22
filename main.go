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
	//customer crud
	router.HandleFunc("/addCustomer", api.AddCustomer).Methods("POST")
	router.HandleFunc("/getAllCustomers", api.GetAllCustomers).Methods("GET")
	router.HandleFunc("/getCustomer/{id}", api.GetCustomerById).Methods("GET")
	router.HandleFunc("/updateCustomer", api.UpdateCustomerData).Methods("PUT")
	router.HandleFunc("/deleteCustomer/{id}", api.DeleteCustomerData).Methods("DELETE")

	//customer +
	router.HandleFunc("/customers/getDueDetails", api.GetAllPaymentPendingCustomers).Methods("GET")

	//order crud
	router.HandleFunc("/addOrder", api.AddOrder).Methods("POST")
	router.HandleFunc("/getAllOrders", api.GetAllOrders).Methods("GET")
	router.HandleFunc("/getOrder/{id}", api.GetOrderById).Methods("GET")
	router.HandleFunc("/updateOrder", api.UpdateOrderData).Methods("PUT")
	router.HandleFunc("/deleteOrder/{id}", api.DeleteOrderData).Methods("DELETE")

	//order +
	// router.HandleFunc("/orders/getLatestOrderId", api.GetLatestOrderId).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	})

	log.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", c.Handler(router)))
}
