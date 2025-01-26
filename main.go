package main

import (
 
	"net/http"
	"todo-app/controllers"
	"todo-app/db"

	"github.com/gorilla/mux"
)

func main() {
	db.InitDB()
	 
	r := mux.NewRouter()

	r.HandleFunc("/register", controllers.Register).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	http.ListenAndServe(":8080", r)
}
 
/* package main

import (
	"net/http"

	"todo-app/controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	 
	r.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", controllers.CreateTask).Methods("POST")

	 
	http.ListenAndServe(":8080", r)
}
 */