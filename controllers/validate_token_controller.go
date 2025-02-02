package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"todo-app/utils"
)

func ValidateToken(w http.ResponseWriter, r *http.Request) {
 
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") 
	w.Header().Set("Access-Control-Allow-Credentials", "true")           
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")   
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == "OPTIONS" {
		log.Println("OPTIONS-запрос, отправляем 200 OK")
		w.WriteHeader(http.StatusOK)
		return
	}

	log.Printf("Получен запрос для проверки токенов: %s", r.URL)
 
	log.Println("🔹 Все заголовки запроса:")
	for name, values := range r.Header {
		log.Printf("%s: %s", name, values)
	}

	log.Println("🔹 Полученные куки:")
	for _, cookie := range r.Cookies() {
		log.Printf("Кука: %s = %s", cookie.Name, cookie.Value)
	}
 
	accessCookie, errAccess := r.Cookie("access_token")
	refreshCookie, errRefresh := r.Cookie("refresh_token")

	if errAccess != nil {
		log.Printf("[WARNING] Нет access_token: %s", errAccess)
	} else {
		log.Printf("[INFO] Найден access_token: %s", accessCookie.Value)
	}

	if errRefresh != nil {
		log.Printf("[WARNING] Нет refresh_token: %s", errRefresh)
	} else {
		log.Printf("[INFO] Найден refresh_token: %s", refreshCookie.Value)
	}
 
	if errAccess != nil && errRefresh != nil {
		log.Println("[ERROR] Оба токена отсутствуют, отклоняем запрос")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"isAuthorized": false,
			"message":      "Session is inactive, please login again",
		})
		return
	}
 
	if errAccess == nil {
		accessToken := accessCookie.Value
		valid, err := utils.CheckAccessToken(accessToken)
		if err != nil {
			log.Printf("[ERROR] Ошибка при проверке access_token: %s", err)
		}
		if valid {
			log.Printf("[SUCCESS] Access token валиден: %s", accessToken)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"isAuthorized": true,
				"message":      "Session is active",
			})
			return
		} else {
			log.Println("[WARNING] Access token невалиден или истек")
		}
	}
 
	if errRefresh == nil {
		refreshToken := refreshCookie.Value
		log.Println("[INFO] Проверяем валидность refresh_token...")

		valid, _ := utils.CheckRefreshToken(refreshToken)
		if !valid {
			log.Println("[ERROR] Refresh token невалиден или истек")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"isAuthorized": false,
				"message":      "Refresh token is invalid or expired",
			})
			return
		}

		log.Println("[SUCCESS] Refresh token валиден, создаем новый access_token...")
		newAccessToken, err := utils.GenerateNewAccessToken("123")
		if err != nil {
			log.Printf("[ERROR] Ошибка при создании нового access_token: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"isAuthorized": false,
				"message":      "Failed to generate a new access token",
			})
			return
		}

		 
		http.SetCookie(w, &http.Cookie{
			Name:     "access_token",
			Value:    newAccessToken,
			HttpOnly: true,
			Path:     "/",
			Expires:  time.Now().Add(1200 * time.Minute),
			Secure:   false, 
			SameSite: http.SameSiteStrictMode,
		})

		log.Println("[INFO] Новый access_token успешно установлен в cookie")

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"isAuthorized": true,
			"message":      "Access token refreshed",
		})
		return
	}

	log.Println("[ERROR] Все токены истекли или отсутствуют, пользователь не авторизован")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"isAuthorized": false,
		"message":      "Session is inactive, please login again",
	})
}
 
func getCookieValue(r *http.Request) (string, error) {
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		fmt.Println("Ошибка при получении cookie:", err)
		return "", fmt.Errorf("error getting cookie: %w", err)
	}

	fmt.Println("Значение refresh_token:", cookie.Value)
	return cookie.Value, nil
}



 
func CookieHandler(w http.ResponseWriter, r *http.Request) {
    
    cookie, err := r.Cookie("refresh_token")
    cookieAccess, errr := r.Cookie("access_token")
    if err != nil {

        fmt.Println("errrrrr")
        if err == http.ErrNoCookie {
            http.Error(w, "Cookie not found", http.StatusUnauthorized)
        
        }
        http.Error(w, "Error retrieving cookie", http.StatusInternalServerError)
       
    }
    fmt.Println(errr)
   
    fmt.Println(w, "Secure Cookierrrr Value: %s", cookie.Value, "ACESS", cookieAccess.Value)
}
