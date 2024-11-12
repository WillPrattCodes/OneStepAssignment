package services

import (
	"backend/internal/database"
	"backend/internal/models"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//find if user exists and compare pass credentials
func AuthenticateUser(email, password string) (models.User, error) {
	//define user struct
	var user models.User
	query := `SELECT id, username, password, email FROM users WHERE email = ?`
	//query db for user
	err := database.DB.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return user, fmt.Errorf("user not found")
	}

	//compare pass and hashed pass
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, fmt.Errorf("invalid password")
	}
	//return user if no err
	return user, nil
}
