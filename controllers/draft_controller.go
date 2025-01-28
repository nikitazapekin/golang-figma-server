package controllers

import (
	"encoding/json"
 
	"net/http"
	"todo-app/db"
	"todo-app/models"
	"strconv"
)

 
func CreateDraft(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		AuthorID    int    `json:"author_id"`
	}

	 
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Invalid input",
		})
		return
	}
 
	err := models.CreateDraft(db.DB, req.Name, req.Description, req.AuthorID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Failed to create draft",
		})
		return
	}
 
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "Draft created successfully",
	})
}
 
func GetAllDrafts(w http.ResponseWriter, r *http.Request) {
	drafts, err := models.GetAllDrafts(db.DB)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Failed to retrieve drafts",
		})
		return
	}
 
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(drafts)
}
func GetDraftByID(w http.ResponseWriter, r *http.Request) {
	 
	draftIDStr := r.URL.Query().Get("draft_id")
 
	draftID, err := strconv.Atoi(draftIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Invalid draft_id parameter",
		})
		return
	}

	 
	draft, err := models.GetDraftByID(db.DB, draftID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Draft not found",
		})
		return
	}

	 
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(draft)
}


func UpdateDraft(w http.ResponseWriter, r *http.Request) {
	var req struct {
		DraftID     int    `json:"draft_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}

 
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Invalid input",
		})
		return
	}
 
	err := models.UpdateDraft(db.DB, req.DraftID, req.Name, req.Description)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Failed to update draft",
		})
		return
	}

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Draft updated successfully",
	})
}
func DeleteDraft(w http.ResponseWriter, r *http.Request) {
	draftIDStr := r.URL.Query().Get("draft_id")

	 
	draftID, err := strconv.Atoi(draftIDStr)
	if err != nil {
	 
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Invalid draft_id parameter",
		})
		return
	}
 
	err = models.DeleteDraft(db.DB, draftID)
	if err != nil {
	 
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Failed to delete draft",
		})
		return
	}

	 
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Draft deleted successfully",
	})
}
