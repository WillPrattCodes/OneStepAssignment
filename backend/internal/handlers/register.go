package handlers

import (
	"backend/internal/models"
	"backend/internal/services"
	"backend/internal/utils"
	"database/sql"
	"encoding/json"
	"net/http"
)

//handle user registration
func RegisterUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
			//declare user struct
			var user models.User
			//decode request body into user struct
			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
					utils.SendErrorResponse(w, "Invalid input", http.StatusBadRequest)
					return
			}

			//call service to register the user
			err = services.RegisterUser(db, user)
			if err != nil {
					switch err {
					case services.ErrInvalidEmail:
							utils.SendErrorResponse(w, "Invalid email format", http.StatusBadRequest)
					case services.ErrEmailAlreadyExists:
							utils.SendErrorResponse(w, "Email already exists", http.StatusConflict)
					case services.ErrUsernameAlreadyExists:
							utils.SendErrorResponse(w, "Username already exists", http.StatusConflict)
					default:
							utils.SendErrorResponse(w, "User registration failed", http.StatusInternalServerError)
					}
					return
			}

			//generate token for the registered user
			token, err := utils.GenerateJWT(user.ID)
			if err != nil {
					utils.SendErrorResponse(w, "Failed to generate token", http.StatusInternalServerError)
					return
			}

			//define response
			response := map[string]interface{}{
				"token": token,
				"user": map[string]interface{}{
						"id":       user.ID,
						"username": user.Username,
						"email":    user.Email,
				},
		}

			//return token in the response
			utils.SendJSONResponse(w, response)
	}
}
