package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func handleGetTest(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Route", r.URL.Path)

	fmt.Println("Route 2", chi.RouteContext(r.Context()).RoutePattern())
	return writeJson(w, http.StatusOK, "test")
}
