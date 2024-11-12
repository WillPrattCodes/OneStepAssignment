package services

import (
	"backend/internal/database"
	"backend/internal/models"
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

//custom err types
var (
	ErrInvalidEmail = fmt.Errorf("invalid email format")
	ErrEmailAlreadyExists = fmt.Errorf("email already exists")
	ErrHashingPassword = fmt.Errorf("failed to hash password")
)

//make sure email is in correct format
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

//helper func to hash password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
};

//register user in db
func RegisterUser(user models.User) error {
	fmt.Println("In RegisterUser")

	//validate email value
	if !isValidEmail(user.Email) {
		return ErrInvalidEmail
	};

	//hash pass
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return ErrHashingPassword
	};

	//insert user into db
	insertQuery := `INSERT INTO users (username, password, email) VALUES (?, ?, ?)`
	_, err = database.DB.Exec(insertQuery, user.Username, hashedPassword, user.Email)
	if err != nil {
		//handle email existing error
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return ErrEmailAlreadyExists
		}
		return fmt.Errorf("failed to register user: %v", err)
	};

	return nil
}
