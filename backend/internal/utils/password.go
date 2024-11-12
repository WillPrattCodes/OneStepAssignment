package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//util function to has user pass
func HashPassword(password string) (string, error) {
	//
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedPass), nil
}