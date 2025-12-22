package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vinit-jpl/restaurant-management/auth-service/internal/dto"
)

// func writeJSON(w http.ResponseWriter, status int, response dto.ApiResponse) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(status)
// 	json.NewEncoder(w).Encode(response)
// }

// func RespondSuccess(w http.ResponseWriter, code int, message string, data interface{}) {
// 	resp := dto.ApiResponse{
// 		Status:  "success",
// 		Code:    code,
// 		Message: message,
// 		Data:    data,
// 	}
// 	writeJSON(w, code, resp)
// }

// func RespondError(w http.ResponseWriter, code int, err error) {
// 	resp := dto.ApiResponse{
// 		Status:  "failed",
// 		Code:    code,
// 		Message: err.Error(),
// 	}
// 	writeJSON(w, code, resp)
// }

func RespondSuccess(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	resp := dto.ApiResponse{
		Status:  "success",
		Code:    statusCode,
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}

func RespondError(w http.ResponseWriter, statusCode int, message string, err error) {
	resp := dto.ApiResponse{
		Status:  "failed",
		Code:    statusCode,
		Message: fmt.Sprintf("%s: %v", message, err),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}
