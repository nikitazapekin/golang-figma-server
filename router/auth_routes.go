
package router

import (
    "todo-app/controllers" 
     "github.com/gorilla/mux"
 
)

 
func SetAuthRoutes(r *mux.Router) {
    r.HandleFunc("/register", controllers.Register).Methods("POST")
    r.HandleFunc("/login", controllers.Login).Methods("POST")
    r.HandleFunc("/validate-token", controllers.CookieHandler).Methods("GET")
}


 