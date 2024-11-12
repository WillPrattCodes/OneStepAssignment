package handlers

import (
	"backend/internal/services"
	"backend/internal/utils"
	"encoding/json"
	"net/http"
)

//handle user log in
func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	//define credentials struct
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	//parse req body into crendtials struct
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		utils.SendErrorResponse(w, "Invalid input", http.StatusBadRequest)
		return
	}

	//authenticate user usering auth service
	user, err := services.AuthenticateUser(credentials.Email, credentials.Password)
	if err != nil {
		utils.SendErrorResponse(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	//user jwt util to generate token
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		utils.SendErrorResponse(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	//define response
	response := map[string]string{
		"token": token,
	}
	//send response
	utils.SendJSONResponse(w, response)
}
