package handlers

import (
	"backend/internal/models"
	"backend/internal/services"
	"backend/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

//handle user registration
func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In RegisterUserHandler")

	//decode req body into user struct
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	fmt.Println(user);
	if err != nil {
		utils.SendErrorResponse(w, "Invalid input", http.StatusBadRequest)
		return
	}

	//call service to register user in db
	err = services.RegisterUser(user)
	if err != nil {
		//error handling 
		switch err {
		case services.ErrInvalidEmail:
			utils.SendErrorResponse(w, "Invalid email format", http.StatusBadRequest)
		case services.ErrEmailAlreadyExists:
			utils.SendErrorResponse(w, "Email already exists", http.StatusConflict)
		case services.ErrHashingPassword:
			utils.SendErrorResponse(w, "Failed to hash password", http.StatusInternalServerError)
		default:
			utils.SendErrorResponse(w, "User registration failed", http.StatusInternalServerError)
		}
		return
	}

	//success message
	utils.SendJSONResponse(w, "User registered successfully")
}
