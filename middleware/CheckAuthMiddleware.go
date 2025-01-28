package middleware

import (
	"net/http"
	"strings"
	"todo-app/utils"
	"encoding/json"
	
)

// CheckAuthMiddleware проверяет токен в заголовке Authorization
func CheckAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Проверяем наличие Authorization заголовка
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "Authorization header is missing",
			})
			return
		}

		// Проверяем правильность формата Bearer токена
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "Invalid token format",
			})
			return
		}

		// Проверяем действительность токена
		valid, err := utils.CheckAccessToken(tokenString)
		if err != nil || !valid {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"message": "Invalid or expired token",
			})
			return
		}

		// Если токен валиден, передаем управление дальше
		next.ServeHTTP(w, r)
	}
}
