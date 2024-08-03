package router

import (
	"encoding/json"
	"net/http"
)

type ApiHandlerFunc func(w http.ResponseWriter, r *http.Request) error

type ApiError struct {
	Error string `json:"error"`
}

func makeHandlerFunc(handler ApiHandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := handler(w, r); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&ApiError{
				Error: err.Error(),
			})
		}
	}
}
