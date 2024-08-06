package router

import (
	"encoding/json"
	"net/http"
)

func AccessDeniedResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(&ApiError{
		Error: "access denied",
	})
}

func SomethingWentWrongResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(&ApiError{
		Error: "something went wrong",
	})
}

func writeJson(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return err
	}
	return nil
}

func writeError(w http.ResponseWriter, status int, message string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(&ApiError{
		Error: message,
	}); err != nil {
		return err
	}
	return nil
}
