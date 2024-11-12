package services

import (
	"backend/internal/database"
	"backend/internal/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// AuthenticateUser validates the user's credentials and returns an error if they are invalid
func AuthenticateUser(email, password string) (models.User, error) {
	// Retrieve the user from the database
	var user models.User
	query := `SELECT id, username, password, email FROM users WHERE email = ?`
	err := database.DB.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return user, fmt.Errorf("user not found")
	}

	// Compare the hashed password with the one provided by the user
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, fmt.Errorf("invalid password")
	}

	return user, nil
}
