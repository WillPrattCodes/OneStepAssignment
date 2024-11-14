package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

//generate jwt token
func GenerateJWT(userID int) (string, error) {
	//define claims
	claims := &jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),//token expiration time
	}
	//define token claims and signature method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//sign token using jwt secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}
	return tokenString, nil
}

//validate jwt token
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	//define claims struct
	claims := jwt.MapClaims{}
	//parse token with claims and jwt secret key
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	//return user info from token
	return claims, nil
}
