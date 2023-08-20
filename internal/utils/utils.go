package utils

import (
	"encoding/json"
	"net/http"
)

type message struct {
	Message string `json:"message"`
}

func ErrorJSON(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(message{msg})
}
