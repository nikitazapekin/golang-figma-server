package views

import (
	"encoding/json"
	"net/http"
)

// RespondJSON writes JSON response to the client
func RespondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
