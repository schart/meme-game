package utils

import (
	"encoding/json"
	"net/http"
	types "shared-library/types"
)

func NewResponse(ok bool, status int, message string) types.Response {
	response := types.Response{
		OK:      ok,
		Status:  status,
		Message: message,
	}

	return response
}

func HandleError(w http.ResponseWriter, statusCode int, errorMessage string) {
	w.Header().Set("Content-Type", "application/json")
	response := NewResponse(false, statusCode, errorMessage)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func HandleSuccess(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	response := NewResponse(true, http.StatusOK, "Successfuly process ")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
