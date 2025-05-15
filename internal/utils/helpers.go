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
	RespondWithJSON(w, status, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, status int, payload any) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)

	return nil
}
