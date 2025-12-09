package utils

import (
	"encoding/json"
	"net/http"

	"github.com/vinit-jpl/restaurant-management/auth-service/internal/dto"
)

func writeJSON(w http.ResponseWriter, status int, response dto.ApiResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func RespondSuccess(w http.ResponseWriter, code int, message string, data interface{}) {
	resp := dto.ApiResponse{
		Status:  "success",
		Code:    code,
		Message: message,
		Data:    data,
	}
	writeJSON(w, code, resp)
}

func RespondError(w http.ResponseWriter, code int, err error) {
	resp := dto.ApiResponse{
		Status:  "failed",
		Code:    code,
		Message: err.Error(),
	}
	writeJSON(w, code, resp)
}
