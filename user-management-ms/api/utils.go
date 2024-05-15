package api

import (
	"encoding/json"
	"net/http"
)

// A set of utility/helper functions that will be reused within the function.
func respondJSON(w http.ResponseWriter, data interface{}, status int) {
	resp_data, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(resp_data))
}
