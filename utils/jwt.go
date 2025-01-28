package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Секретный ключ для подписывания токенов
var JwtSecret = []byte("your_secret_key")
func GenerateJWT(userID int) (string, string, error) {
	// Генерация access_token
	accessTokenClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Minute * 15).Unix(), // Время жизни 15 минут
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString(JwtSecret)
	if err != nil {
		return "", "", err
	}

	// Генерация refresh_token
	refreshTokenClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // Время жизни 7 дней
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString(JwtSecret)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}


func CheckRefreshToken(refreshToken string) (bool, int) {
	// Парсинг refresh_token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return JwtSecret, nil
	})

	if err != nil || !token.Valid {
		return false, 0
	}

	// Извлекаем user_id из claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["user_id"] == nil {
		return false, 0
	}

	userID := int(claims["user_id"].(float64)) // Преобразуем в int
	return true, userID
}

func CheckAccessToken(tokenString string) (bool, error) {
	// Парсинг access_token
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