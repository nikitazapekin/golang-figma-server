package controllers

import (
	"encoding/json"
	"net/http"
	"todo-app/utils"
	// "strings"
	// "github.com/dgrijalva/jwt-go"
)
// Функция для проверки токенов
func ValidateToken(w http.ResponseWriter, r *http.Request) {
    // Получаем refresh_token из cookies
    tokenCookie, err := r.Cookie("refresh_token")
    if err != nil || tokenCookie.Value == "" {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(map[string]interface{}{"valid": false})
        return
    }

    // Если Authorization header пустой, проверяем refresh_token
    if r.Header.Get("Authorization") == "" {
        // Проверяем refresh_token
        valid, _ := utils.CheckRefreshToken(tokenCookie.Value)
        if !valid {
            w.WriteHeader(http.StatusUnauthorized)
            json.NewEncoder(w).Encode(map[string]interface{}{"valid": false})
            return
        }

        // Если refresh_token действителен, генерируем новый access_token
        newAccessToken, _, err := utils.GenerateJWT(123) // Замените 123 на реальный userID
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(map[string]interface{}{"error": "unable to generate access token"})
            return
        }

        json.NewEncoder(w).Encode(map[string]interface{}{"valid": true, "access_token": newAccessToken})
    } else {
        // Проверяем access_token
        tokenString := r.Header.Get("Authorization")
        valid, _ := utils.CheckAccessToken(tokenString)
        if !valid {
            w.WriteHeader(http.StatusUnauthorized)
            json.NewEncoder(w).Encode(map[string]interface{}{"valid": false})
            return
        }
        json.NewEncoder(w).Encode(map[string]interface{}{"valid": true})
    }
}
