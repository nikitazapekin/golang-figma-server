package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-app/db"
	"todo-app/router"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
    db.InitDB()

    r := mux.NewRouter()
    router.InitRoutes(r)

    // Добавляем обработчик для получения cookie
    r.HandleFunc("/get-cookie", cookieHandler).Methods("GET")

    corsHandler := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000"}, // Разрешаем только фронтенд
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Content-Type", "Authorization"},
        AllowCredentials: true, // ОБЯЗАТЕЛЬНО, чтобы браузер отправлял куки
    })

    handler := corsHandler.Handler(r)

    log.Println("🚀 Сервер запущен на :8080")
    err := http.ListenAndServe(":8080", handler)
    if err != nil {
        log.Fatal("Ошибка запуска сервера:", err)
    }
}
 
// Переименованный обработчик
func cookieHandler(w http.ResponseWriter, r *http.Request) {
    // Получаем куку по имени
    cookie, err := r.Cookie("refresh_token")
    cookieAccess, errr := r.Cookie("access_token")
    if err != nil {

        fmt.Println("errrrrr")
        if err == http.ErrNoCookie {
            http.Error(w, "Cookie not found", http.StatusUnauthorized)
            return
        }
        http.Error(w, "Error retrieving cookie", http.StatusInternalServerError)
        return
    }
    fmt.Println(errr)
    // Выводим значение куки
    fmt.Println(w, "Secure Cookierrrr Value: %s", cookie.Value, "ACESS", cookieAccess.Value)
}
 