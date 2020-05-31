package handler

import (
	"encoding/json"
	"net/http"
)

// Response is the struct which will be sent over API response
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func writeServerErrorIfErrNotNil(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
}

func writeSuccessResponse(w http.ResponseWriter, data interface{}) {
	res := Response{
		Success: true,
		Data:    data,
	}

	body, err := json.Marshal(res)
	writeServerErrorIfErrNotNil(w, err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(body)
	writeServerErrorIfErrNotNil(w, err)
}

func writeErrorResponse(w http.ResponseWriter, targetErr string, status int) {
	res := Response{
		Success: false,
		Error:   targetErr,
	}

	body, err := json.Marshal(res)
	writeServerErrorIfErrNotNil(w, err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(body)
	writeServerErrorIfErrNotNil(w, err)
}
