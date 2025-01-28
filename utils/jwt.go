package utils

import (
	"encoding/json"
	"errors"
	"net/http"
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
 


// Пример endpoint для обновления access_token
func RefreshToken(w http.ResponseWriter, r *http.Request) {
    // Получаем refresh_token из cookie
    cookie, err := r.Cookie("refresh_token")
    if err != nil {
        http.Error(w, "Refresh token not found", http.StatusUnauthorized)
        return
    }

    // Проверяем refresh_token и генерируем новый access_token
    accessToken, err := generateNewAccessToken(cookie.Value)
    if err != nil {
        http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
        return
    }

    // Отправляем новый access_token клиенту
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "access_token": accessToken,
    })
}


// Функция для генерации нового access_token
func generateNewAccessToken(refreshToken string) (string, error) {
	// Проверяем действительность refresh_token (например, проверка в базе данных или верификация JWT)
	claims := &jwt.StandardClaims{}
	_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	if err != nil {
		return "", err // Если ошибка верификации, возвращаем ошибку
	}

	// Если refresh_token действителен, генерируем новый access_token
	newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":     time.Now().Add(time.Hour * 1).Unix(), // Новый access_token с временем жизни 1 час
		"user_id": claims.Subject, // Здесь предполагается, что user_id сохранен в claims
	})

	// Подписываем новый токен
	tokenString, err := newAccessToken.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}