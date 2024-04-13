package handlers

import (
	"encoding/json"
	"net/http"
)

type ErrorResponseHandler struct {
	Message string `json:"message"`
}

func JSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func ErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	errResponse := ErrorResponseHandler{
		Message: message,
	}

	JSONResponse(w, statusCode, errResponse)
}
