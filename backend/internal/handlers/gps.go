package handlers

import (
	"backend/internal/services"
	"backend/internal/utils"
	"net/http"
)

//handler to request data from gps service
func GetGPSDataHandler(w http.ResponseWriter, r *http.Request) {
	//invoke fetch gps service function
	data, err := services.FetchGPSData()
	if err != nil {
		utils.SendErrorResponse(w, "Failed to get GPS data", http.StatusInternalServerError)
		return
	}

	//send back response
	utils.SendJSONResponse(w, data)
}
