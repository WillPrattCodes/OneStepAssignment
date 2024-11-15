package services

import (
	"backend/internal/models"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//find if user exists and compare pass credentials
func AuthenticateUser(db *sql.DB, email, password string) (models.User, error) {
	//define user struct
	var user models.User
	query := `SELECT id, username, password, email FROM users WHERE email = ?`
	//query db for user
	err := db.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
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
