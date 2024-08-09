package router

import (
	"encoding/json"
	"net/http"
)

func getRequestBody[T any](r *http.Request) (*T, error) {
	var body T
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		return nil, err
	}
	return &body, nil
}
