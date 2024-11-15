package handlers

import (
	"backend/internal/models"
	"backend/internal/services"
	"backend/internal/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

//get user preferences handler
func GetPreferencesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//extract userID from context
		userID, ok := r.Context().Value(utils.UserIDKey).(int)
		if !ok {
			utils.SendErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		//fetch user preferences
		prefs, err := services.FetchPreferences(db, userID)
		if err != nil {
			utils.SendErrorResponse(w, "Failed to fetch preferences", http.StatusInternalServerError)
			return
		}
		//send response
		utils.SendJSONResponse(w, prefs)
	}
}

//save or update user preferences handler
func SetPreferencesHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//extract userID from context
		userIDValue := r.Context().Value(utils.UserIDKey)

		//check if the value is nil
		if userIDValue == nil {
			http.Error(w, "unauthorized: user ID not found in context", http.StatusUnauthorized)
			return
		}

		//typer assert userID to int
		userID, ok := userIDValue.(int)
		if !ok {
			fmt.Println("type assertion failed for userID")
			http.Error(w, "invalid user ID in context", http.StatusUnauthorized)
			return
		}

		//create preferences object
		var prefs models.Preferences
		//parse request body
		err := json.NewDecoder(r.Body).Decode(&prefs)
		if err != nil {
			utils.SendErrorResponse(w, "Invalid input", http.StatusBadRequest)
			return
		}

		//use service to save or update preferences
		err = services.SaveOrUpdatePreferences(db, userID, prefs)
		if err != nil {
			utils.SendErrorResponse(w, "Failed to save preferences", http.StatusInternalServerError)
			return
		}
		//send response
		utils.SendJSONResponse(w, "Preferences saved successfully")
	}
}
