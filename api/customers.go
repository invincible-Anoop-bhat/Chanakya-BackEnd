package api

import (
	"Chanakya-BackEnd/model"
	"encoding/json"
	"net/http"
)

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	dbdata := GetCustomersFromDB()
	data := model.CopyArrayToCustomer(dbdata)
	respondJSON(w, http.StatusOK, data)
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"hello\": \"world\"}"))
}

func respondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
