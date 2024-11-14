package services

import (
	"backend/internal/models"
	"backend/internal/utils"
	"database/sql"
	"fmt"
	"regexp"
)

//custom err types
var (
	ErrInvalidEmail = fmt.Errorf("invalid email format")
	ErrEmailAlreadyExists = fmt.Errorf("email already exists")
	ErrUsernameAlreadyExists = fmt.Errorf("username already exists")
	ErrHashingPassword = fmt.Errorf("failed to hash password")
)

//make sure email is in correct format
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

//register user in db
func RegisterUser(db *sql.DB, user models.User) error {

	//validate email value
	if !isValidEmail(user.Email) {
		return ErrInvalidEmail
	}

	//check if email already exists
	var existingEmail string
	err := db.QueryRow("SELECT email FROM users WHERE email = ?", user.Email).Scan(&existingEmail)
	if err == nil {
		return ErrEmailAlreadyExists
	} else if err != sql.ErrNoRows {
		return fmt.Errorf("database error while checking email: %v", err)
	}

	//check if username already exists
	var existingUsername string
	err = db.QueryRow("SELECT username FROM users WHERE username = ?", user.Username).Scan(&existingUsername)
	if err == nil {
		return ErrUsernameAlreadyExists
	} else if err != sql.ErrNoRows {
		return fmt.Errorf("database error while checking username: %v", err)
	}

	//hash pass
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return ErrHashingPassword
	}

	//insert user into db
	insertQuery := `INSERT INTO users (username, password, email) VALUES (?, ?, ?)`
	_, err = db.Exec(insertQuery, user.Username, hashedPassword, user.Email)
	if err != nil {
		return fmt.Errorf("failed to register user: %v", err)
	}

	return nil
}
