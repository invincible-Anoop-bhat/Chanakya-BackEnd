package api

import (
	"Chanakya-BackEnd/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Get ALL Customers
func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	dbdata := model.GetCustomersFromDB()
	data := model.CopyArrayToCustomer(dbdata)
	respondJSON(w, http.StatusOK, data)
}

//Get Customer By id
func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_string, ok := vars["id"]
	if !ok {
		log.Println("id is missing in parameters")
	}
	id, err := strconv.Atoi(id_string)
	if err != nil {
		http.Error(w, "Improper value of customer Id", http.StatusBadRequest)
		return
	}
	dbdata := model.GetCustomerbyIDFromDB(id)
	data := dbdata.CopyToCustomer()
	respondJSON(w, http.StatusOK, data)
}

//INSERT ONE Customer data (select by id)
func AddCustomer(w http.ResponseWriter, r *http.Request) {
	requestBody := model.Customer{}
	json.NewDecoder(r.Body).Decode(&requestBody)
	dataToInsert := model.CopyToCustomerDB(requestBody)
	if len(dataToInsert.CName) == 0 {
		//this is invalid data, so it should not
		http.Error(w, "Required field not found : customer name", http.StatusBadRequest)
		return
	}
	//Checks if customer with same name already exists in DB
	exists, err := model.CheckCustomerExists(dataToInsert.CName)
	// log.Println("Exists : ", exists)
	if err != nil {
		// log.Println("Error checking customer duplication : ", err.Error())
		http.Error(w, "Error checking customer duplication.", http.StatusInternalServerError)
		return
	} else if exists {
		http.Error(w, "Customer already exists", http.StatusBadRequest)
		return
	}

	err = model.InsertCustomerToDB(dataToInsert)
	if err != nil {
		http.Error(w, "Error Adding Customer details", http.StatusInternalServerError)
		return
	}
	respondJSON(w, http.StatusCreated, requestBody)
}

//UPDATE ONE Customer (Select by id)
func UpdateCustomerData(w http.ResponseWriter, r *http.Request) {
	requestBody := model.Customer{}
	json.NewDecoder(r.Body).Decode(&requestBody)
	dataToUpdate := model.CopyToCustomerDB(requestBody)
	err := model.UpdateCustomerInDB(dataToUpdate)
	if err != nil {
		http.Error(w, "Error Adding Customer details", http.StatusInternalServerError)
	}
	respondJSON(w, http.StatusCreated, requestBody)
}

//DELETE ONE Customer (select by id)
func DeleteCustomerData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_string, ok := vars["id"]
	if !ok {
		log.Println("id is missing in parameters")
	}
	CustomerId, err := strconv.Atoi(id_string)
	if err != nil {
		http.Error(w, "Improper value of customer Id", http.StatusBadRequest)
		return
	}
	model.DeleteCustomerFromDB(CustomerId)
	respondJSON(w, http.StatusOK, "Success")
}

func GetAllPaymentPendingCustomers(w http.ResponseWriter, r *http.Request) {
	Customers := model.GetAllPaymentPendingCustomersFromDB()
	respondJSON(w, http.StatusOK, Customers)
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"hello\": \"world\"}"))
}

func respondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "http://localho st:4200")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
