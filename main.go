
package main

import (
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
 
    corsHandler := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000"},   
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},  
        AllowedHeaders:   []string{"Content-Type", "Authorization"},  
        AllowCredentials: true,  
    })

 
    handler := corsHandler.Handler(r)
 
    http.ListenAndServe(":8080", handler)
}

/*
package main

import (
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

    corsHandler := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:3000"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
        AllowedHeaders:   []string{"Content-Type"},
        AllowCredentials: true,
    })

    handler := corsHandler.Handler(r)

    http.ListenAndServe(":8080", handler)
}

 */