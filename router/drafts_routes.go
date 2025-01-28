package router

import (
	"net/http"
	"todo-app/controllers"
	"todo-app/middleware"

	"github.com/gorilla/mux"
)

func SetDraftsRoutes(r *mux.Router) {
	r.HandleFunc("/drafts", middleware.CheckAuthMiddleware(http.HandlerFunc(controllers.GetAllDrafts))).Methods("GET")
	r.HandleFunc("/drafts/{id:[0-9]+}", middleware.CheckAuthMiddleware(http.HandlerFunc(controllers.GetDraftByID))).Methods("GET")
	r.HandleFunc("/createDraft", middleware.CheckAuthMiddleware(http.HandlerFunc(controllers.CreateDraft))).Methods("POST")
	r.HandleFunc("/drafts", middleware.CheckAuthMiddleware(http.HandlerFunc(controllers.UpdateDraft))).Methods("PUT")
	r.HandleFunc("/drafts", middleware.CheckAuthMiddleware(http.HandlerFunc(controllers.DeleteDraft))).Methods("DELETE")
}

