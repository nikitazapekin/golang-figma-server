package router

import (
	"github.com/gorilla/mux"
	"todo-app/controllers"
)

func SetDraftsRoutes(r *mux.Router) {
	r.HandleFunc("/drafts", controllers.GetAllDrafts).Methods("GET")
	r.HandleFunc("/drafts/{id:[0-9]+}", controllers.GetDraftByID).Methods("GET")
	r.HandleFunc("/createDraft", controllers.CreateDraft).Methods("POST")
	r.HandleFunc("/drafts", controllers.UpdateDraft).Methods("PUT")
	r.HandleFunc("/drafts", controllers.DeleteDraft).Methods("DELETE")
}
