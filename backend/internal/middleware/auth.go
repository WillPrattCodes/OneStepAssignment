package middleware

import (
	"backend/internal/utils"
	"context"
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
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		//extract userID from claims
		userID, ok := claims["userID"].(float64)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}
		//add userID to req context
		ctx := context.WithValue(r.Context(), utils.UserIDKey, int(userID))

		//call next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
