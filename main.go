 




  package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"todo-app/db"
	"todo-app/router"
	"todo-app/utils"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	db.InitDB()
	r := mux.NewRouter()
	router.InitRoutes(r)

	r.HandleFunc("/get-cookie", cookieHandler).Methods("GET")

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Cookie"},
		AllowCredentials: true,
	})

	handler := corsHandler.Handler(r)

	log.Println("🚀 Сервер запущен на :8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal("Ошибка запуска сервера:", err)
	}
}

func cookieHandler(w http.ResponseWriter, r *http.Request) {
	accessCookie, errAccess := r.Cookie("access_token")
	refreshCookie, errRefresh := r.Cookie("refresh_token")

	if errAccess != nil {
		log.Println("[WARNING] Access token отсутствует или истек")
	} else {
		valid, _ := utils.CheckAccessToken(accessCookie.Value)
		if valid {
			log.Println("[INFO] Access token валиден")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintln(w, "Access token валиден")
			return
		}
	}

	if errRefresh != nil {
		log.Println("[ERROR] Refresh token отсутствует или истек")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Session is inactive, please login again")
		return
	}

	valid, _ := utils.CheckRefreshToken(refreshCookie.Value)
	if !valid {
		log.Println("[ERROR] Refresh token невалиден")
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Refresh token is invalid or expired")
		return
	}

	newAccessToken, err := utils.GenerateNewAccessToken("123")
	if err != nil {
		log.Println("[ERROR] Ошибка генерации нового access_token")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Failed to generate a new access token")
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    newAccessToken,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(10 * time.Minute),
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	})

	log.Println("[SUCCESS] Новый access_token установлен")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Access token refreshed")
}
