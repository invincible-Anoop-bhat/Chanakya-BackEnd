package api

import (
	"net/http"
)

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, http.StatusOK, "All Customers")
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"hello\": \"world\"}"))
}

func respondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte("data"))
}
