package router

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func AccessDeniedResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(&ApiError{
		Error: "Access denied",
	})
}

func SomethingWentWrongResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(&ApiError{
		Error: "Something went wrong",
	})
}

func ReturnAccessDenied() error {
	return fmt.Errorf("access denied")
}
func ReturnSomethingWentWrong() error {
	return fmt.Errorf("something went wrong")
}
