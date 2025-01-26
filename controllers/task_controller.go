package controllers

import (
	"encoding/json"
	"net/http"
	"todo-app/models"
	"todo-app/views"
)

// GetTasks handles the GET request to fetch all tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks := models.GetAllTasks()
	views.RespondJSON(w, http.StatusOK, tasks)
}

// CreateTask handles the POST request to create a new task
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		views.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}

	models.AddTask(task)
	views.RespondJSON(w, http.StatusCreated, task)
}
