package main

import (
	"net/http"
	"todo-app/controllers"
	"todo-app/db"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	 
	db.InitDB()

	 
	r := mux.NewRouter()

	 
	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	// Настройка CORS
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, 
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},  
		AllowedHeaders: []string{"Content-Type"}, 
		AllowCredentials: true,  
	})

 
	handler := corsHandler.Handler(r)

 
	http.ListenAndServe(":8080", handler)
}
 