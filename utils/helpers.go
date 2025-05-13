package utils

import (
	"encoding/json"
	"net/http"
)

const (
	ContentType string = "Content-Type"
	ContentJSON string = "application/json"
)

func RespondWithError(w http.ResponseWriter, status int, message string) {
	respondWithJSON(w, status, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, status int, payload any) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}
