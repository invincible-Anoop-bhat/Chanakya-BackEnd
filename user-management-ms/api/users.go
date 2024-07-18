package api

import "net/http"

//CRUD for users

//Create
func createUser(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, "This is create new user", http.StatusOK)
}

//Read
func getUserById(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("This is get user by id"))
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, "This is get all users", http.StatusOK)
}

//Update
func updateUser(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, "This is update an user", http.StatusOK)
}

//Delete
func deleteUserById(w http.ResponseWriter, r *http.Request) {
	respondJSON(w, "This is delete an user", http.StatusOK)
}
