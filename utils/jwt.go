package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)
var JwtSecret = []byte("your_secret_key")
func GenerateJWT(userID int) (string, string, error) {
	accessTokenClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * 15).Unix(),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString(JwtSecret)
	if err != nil {
		return "", "", err
	}

 
	refreshTokenClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),  
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString(JwtSecret)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}


func CheckRefreshToken(refreshToken string) (bool, int) {
	 
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return JwtSecret, nil
	})

	if err != nil || !token.Valid {
		return false, 0
	}
 
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["user_id"] == nil {
		return false, 0
	}

	userID := int(claims["user_id"].(float64)) 
	return true, userID
}

func CheckAccessToken(tokenString string) (bool, error) {
 
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return JwtSecret, nil
	})

	if err != nil || !token.Valid {
		return false, err
	}
	return true, nil
}
  /*
func RefreshToken(w http.ResponseWriter, r *http.Request) {
 
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		http.Error(w, "Refresh token not found", http.StatusUnauthorized)
		return
	}
 
	accessToken, err := generateNewAccessToken(cookie.Value)
	if err != nil {
		http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
		return
	}
 
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"access_token": accessToken,
	})
}
	*/


	func RefreshToken(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("refresh_token")
		if err != nil || cookie.Value == "" {
			// Удаление куки
			http.SetCookie(w, &http.Cookie{
				Name:     "refresh_token",
				Value:    "",
				Path:     "/",
				Expires:  time.Unix(0, 0),
				HttpOnly: true,
				Secure:   false,
				SameSite: http.SameSiteStrictMode,
			})
	
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "Refresh token not found or invalid",
			})
			return
		}
	
		accessToken, err := GenerateNewAccessToken(cookie.Value)
		if err != nil {
			 
			http.SetCookie(w, &http.Cookie{
				Name:     "refresh_token",
				Value:    "",
				Path:     "/",
				Expires:  time.Unix(0, 0),
				HttpOnly: true,
				Secure:   false,
				SameSite: http.SameSiteStrictMode,
			})
	
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": "Invalid refresh token",
			})
			return
		}
	 
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"access_token": accessToken,
		})
	}

	

func GenerateNewAccessToken(refreshToken string) (string, error) {
	valid, userID := CheckRefreshToken(refreshToken)
	if !valid {
		return "", fmt.Errorf("invalid refresh token")
	}
	accessToken, _, err := GenerateJWT(userID)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}


