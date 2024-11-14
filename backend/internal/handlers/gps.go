package handlers

import (
	"backend/internal/services"
	"backend/internal/utils"
	"database/sql"
	"fmt"
	"net/http"
)

//handler to request data from gps service
func GetGPSDataHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		
		//invoke fetch gps service function
		data, err := services.FetchGPSData(db)
		if err != nil {
			utils.SendErrorResponse(w, "Failed to get GPS data", http.StatusInternalServerError)
			return
		}
		fmt.Printf("data: %v\n", data)
		//send back response
		utils.SendJSONResponse(w, data)
	}
}
