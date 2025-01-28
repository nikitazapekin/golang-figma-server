package middleware

import (
	"net/http"
	"strings"
	"todo-app/utils"
	"encoding/json"
	
)

func CheckAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "Authorization header is missing",
			})
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "Invalid token format",
			})
			return
		}
		valid, err := utils.CheckAccessToken(tokenString)
		if err != nil || !valid {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "Invalid or expired token",
			})
			return
		}
		next.ServeHTTP(w, r)
	}
}
