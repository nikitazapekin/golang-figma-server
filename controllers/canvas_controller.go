package controllers

import (
	"encoding/json"
	"net/http"
	"todo-app/db"
	"todo-app/models"
)
 
func SetCanvas(w http.ResponseWriter, r *http.Request) {
	var req models.CanvasRequest

 
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"code": 400, "message": "Invalid input"}`, http.StatusBadRequest)
		return
	}
 
	if req.ID == "" {
		http.Error(w, `{"code": 400, "message": "ID is required"}`, http.StatusBadRequest)
		return
	}
 
	tx, err := db.DB.Begin()
	if err != nil {
		http.Error(w, `{"code": 500, "message": "Database error"}`, http.StatusInternalServerError)
		return
	}
	defer tx.Rollback() 
 
	for _, figure := range req.ArrayOfFigures {
		_, err := tx.Exec(`
			INSERT INTO figures (draft_id, coordX, coordY, type, width, height, opacity, border, background, stroke, strokeColor, shadowColor, shadowX, shadowY, layout)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)`,
			req.ID, figure.CoordX, figure.CoordY, figure.Type, figure.Width, figure.Height,
			figure.Opacity, figure.Border, figure.Background, figure.Stroke, figure.StrokeColor,
			figure.ShadowColor, figure.ShadowX, figure.ShadowY, figure.Layout,
		)
		if err != nil {
			http.Error(w, `{"code": 500, "message": "Failed to insert figure"}`, http.StatusInternalServerError)
			return
		}
	}
	for _, line := range req.ArrayOfLines {
	 
		jsonPath, err := json.Marshal(line.Path)
	   if err != nil {
		   http.Error(w, `{"code": 400, "message": "Invalid path format"}`, http.StatusBadRequest)
		   return
	   }  
		_, err = tx.Exec(`
		INSERT INTO lines (draft_id, coordX, coordY, type, width, height, path, strokeWidth, color, layout)
		VALUES ($1, $2, $3, $4, $5, $6, $7::jsonb, $8, $9, $10)`,
		req.ID, line.CoordX, line.CoordY, line.Type, line.Width, line.Height,
	jsonPath, line.StrokeWidth, line.Color, line.Layout,
	)
	
		if err != nil {
			http.Error(w, `{"code": 500, "message": "Failed to insert line"}`, http.StatusInternalServerError)
			return
		}
	}

	
	if err := tx.Commit(); err != nil {
		http.Error(w, `{"code": 500, "message": "Transaction commit failed"}`, http.StatusInternalServerError)
		return
	}
 
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code":    http.StatusCreated,
		"message": "Canvas data saved successfully",
	})
}
 