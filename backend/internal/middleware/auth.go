package middleware

import (
	"backend/internal/utils"
	"net/http"
	"strings"
)

//ensure req has a valid jwt token
func AuthMiddleware(next http.Handler) http.Handler {
	//use closure to return http.HandlerFunc
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//check if auth header exists
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		//extract token from header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		//validate token using utils
		_, err := utils.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		//call next handler
		next.ServeHTTP(w, r)
	})
}
