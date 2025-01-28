package controllers

import (
	"encoding/json"
	"net/http"
	"todo-app/utils"

)
func ValidateToken(w http.ResponseWriter, r *http.Request) {
    tokenCookie, err := r.Cookie("refresh_token")
    if err != nil || tokenCookie.Value == "" {
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode(map[string]interface{}{"valid": false})
        return
    }

    if r.Header.Get("Authorization") == "" {
        valid, _ := utils.CheckRefreshToken(tokenCookie.Value)
        if !valid {
            w.WriteHeader(http.StatusUnauthorized)
            json.NewEncoder(w).Encode(map[string]interface{}{"valid": false})
            return
        }
        newAccessToken, _, err := utils.GenerateJWT(123) 
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            json.NewEncoder(w).Encode(map[string]interface{}{"error": "unable to generate access token"})
            return
        }

        json.NewEncoder(w).Encode(map[string]interface{}{"valid": true, "access_token": newAccessToken})
    } else {
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
