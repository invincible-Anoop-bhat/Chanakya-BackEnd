package api

import (
	"Chanakya-BackEnd/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Get ALL Orders
func GetAllOrders(w http.ResponseWriter, r *http.Request) {
	dbdata := model.GetOrdersFromDB()
	data := model.CopyArrayToOrder(dbdata)
	respondJSON(w, http.StatusOK, data)
}

//Get Order By id
func GetOrderById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_string, ok := vars["id"]
	if !ok {
		log.Println("id is missing in parameters")
	}
	id, err := strconv.Atoi(id_string)
	if err != nil {
		http.Error(w, "Improper value of order Id", http.StatusBadRequest)
		return
	}
	dbdata := model.GetOrderbyIDFromDB(id)
	data := dbdata.CopyToOrder()
	respondJSON(w, http.StatusOK, data)
}

//INSERT ONE Order data (select by id)
func AddOrder(w http.ResponseWriter, r *http.Request) {
	requestBody := model.Order{}
	json.NewDecoder(r.Body).Decode(&requestBody)
	dataToInsert := model.CopyToOrderDB(requestBody)
	err := model.InsertOrderToDB(dataToInsert)
	if err != nil {
		http.Error(w, "Error Adding Order details", http.StatusInternalServerError)
	}
	respondJSON(w, http.StatusCreated, requestBody)
}

//UPDATE ONE Order (Select by id)
func UpdateOrderData(w http.ResponseWriter, r *http.Request) {
	requestBody := model.Order{}
	json.NewDecoder(r.Body).Decode(&requestBody)
	dataToUpdate := model.CopyToOrderDB(requestBody)
	err := model.UpdateOrderInDB(dataToUpdate)
	if err != nil {
		http.Error(w, "Error Adding Order details", http.StatusInternalServerError)
	}
	respondJSON(w, http.StatusCreated, requestBody)
}

//DELETE ONE Order (select by id)
func DeleteOrderData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id_string, ok := vars["id"]
	if !ok {
		log.Println("id is missing in parameters")
	}
	OrderId, err := strconv.Atoi(id_string)
	if err != nil {
		http.Error(w, "Improper value of order Id", http.StatusBadRequest)
		return
	}
	model.DeleteOrderFromDB(OrderId)
	respondJSON(w, http.StatusOK, "Success")
}
