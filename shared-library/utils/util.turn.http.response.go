package utils

import (
	"net/http"
	types "shared-library/types"

	"github.com/goccy/go-json"
)

func NewResponse(ok bool, status int, message string, params map[string]interface{}) types.Response {
	response := types.Response{
		OK:      ok,
		Status:  status,
		Message: message,
		Params:  params,
	}

	return response
}

func HandleError(w http.ResponseWriter, statusCode int, errorMessage string) {
	w.Header().Set("Content-Type", "application/json")
	response := NewResponse(false, statusCode, errorMessage, map[string]interface{}{})
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func HandleSuccess(w http.ResponseWriter, params  map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	response := NewResponse(true, http.StatusOK, "Successfuly process ", params)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
