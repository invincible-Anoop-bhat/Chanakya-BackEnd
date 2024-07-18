package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Define mapping for routes and their handlers.
func MapHandlersToRoutes(router *mux.Router) {

	router.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to User Management Backend!"))
	}).Methods("GET")

	router.HandleFunc("/allusers", getAllUsers).Methods("GET")
	router.HandleFunc("/user/{id}", getUserById).Methods("GET")
	router.HandleFunc("/createuser", createUser).Methods("POST")
	router.HandleFunc("/updateuser", updateUser).Methods("PUT")
	router.HandleFunc("/removeuser/{id}", deleteUserById).Methods("DELETE")
}
