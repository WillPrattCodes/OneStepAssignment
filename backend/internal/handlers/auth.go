package handlers

import (
	"backend/internal/services"
	"backend/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

//handle user log in
func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	//define credentials struct
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	//decode req body into crendtials struct
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		utils.SendErrorResponse(w, "Invalid input", http.StatusBadRequest)
		return
	}

	//user authentication service to authenticate user
	user, err := services.AuthenticateUser(credentials.Email, credentials.Password)
	if err != nil {
		fmt.Println("Authentication failed:", err)
		utils.SendErrorResponse(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	//success message response
	utils.SendJSONResponse(w, fmt.Sprintf("Welcome, %s!", user.Username))
}
