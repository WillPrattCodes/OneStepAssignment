package utils

import (
	"encoding/json"
	"net/http"
)

//success response function
func SendJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//convert to json
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

//err response function
func SendErrorResponse(w http.ResponseWriter, errorMessage string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{
		"error": errorMessage,
	})
}
