/* package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"todo-app/db"
	"todo-app/models"
	"todo-app/utils"
	
)
 
	func Register(w http.ResponseWriter, r *http.Request) {
 
		var req struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}
	
 
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": "Invalid input",
			})
			return
		}
 
		err := models.CreateUser(db.DB, req.Username, req.Email, req.Password)
		if err != nil {
			fmt.Println("err", err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": "Failed to create user",
			})
			return
		}
	
 
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusCreated,
			"message": "User created successfully",
		})
	}
	
 
func Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := models.AuthenticateUser(db.DB, req.Email, req.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	accessToken, refreshToken, err := utils.GenerateJWT(user.ID)
	if err != nil {
		http.Error(w, "Failed to generate tokens", http.StatusInternalServerError)
		return
	}

 
 
expiration := time.Now().Add(7 * 24 * time.Hour)  

http.SetCookie(w, &http.Cookie{
    Name:     "refresh_token",
    Value:    refreshToken,
    HttpOnly: true,
    Path:     "/",
    Expires:  expiration,  
    Secure:   false, 
    SameSite: http.SameSiteStrictMode,  
})



	json.NewEncoder(w).Encode(map[string]string{"access_token": accessToken})
}


  */

  package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"todo-app/db"
	"todo-app/models"
	"todo-app/utils"
)

// Register - контроллер для регистрации пользователя
func Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Invalid input",
		})
		return
	}

	err := models.CreateUser(db.DB, req.Username, req.Email, req.Password)
	if err != nil {
		fmt.Println("err", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Failed to create user",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "User created successfully",
	})
}

// Login - контроллер для логина пользователя
func Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Аутентификация пользователя
	user, err := models.AuthenticateUser(db.DB, req.Email, req.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Генерация токенов
	accessToken, refreshToken, err := utils.GenerateJWT(user.ID)
	if err != nil {
		http.Error(w, "Failed to generate tokens", http.StatusInternalServerError)
		return
	}
 
	refreshTokenExpiration := time.Now().Add( 2*time.Minute)
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HttpOnly: true, 
		Path:     "/",
		Expires:  refreshTokenExpiration,
		Secure:   false,  
		SameSite: http.SameSiteStrictMode,
	})

 
	accessTokenExpiration := time.Now().Add(1 * time.Minute)  
	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		HttpOnly: true, // Только сервер может получить токен
		Path:     "/",
		Expires:  accessTokenExpiration,
		Secure:   false, // Установите в true, если используете HTTPS
		//SameSite: http.SameSiteStrictMode,
		SameSite: http.SameSiteNoneMode, 
	})

	// Возвращаем access_token в теле ответа (если нужно для клиента)
	json.NewEncoder(w).Encode(map[string]string{
		"access_token": accessToken,
	})
}
