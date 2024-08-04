package router

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kaung-minkhant/go-restaurent/auth"
	"github.com/kaung-minkhant/go-restaurent/database"
	"github.com/kaung-minkhant/go-restaurent/database/models"
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

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("x-access-token")
		if token == "" {
			AccessDeniedResponse(w)
			return
		}

		claims, err := auth.ValidateJWT(token)
		if err != nil {
			AccessDeniedResponse(w)
			return
		}
		permission, err := getPermission(r)
		if err != nil {
			AccessDeniedResponse(w)
			return
		}

		rolePermission, err := getRoleFromPermission(r, permission)
		if err != nil {
			AccessDeniedResponse(w)
			return
		}

		if rolePermission.Role != claims.EmployeeRole {
			AccessDeniedResponse(w)
			return
		}

		ctx := context.WithValue(r.Context(), "ctx-claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getPermission(r *http.Request) (models.Permission, error) {
	routePattern := chi.RouteContext(r.Context()).RoutePattern()
	return database.Db.GetPermissionByMethodAndRoute(r.Context(), models.GetPermissionByMethodAndRouteParams{
		Method: r.Method,
		Route:  routePattern,
	})
}

func getRoleFromPermission(r *http.Request, permission models.Permission) (models.RolePermission, error) {
	return database.Db.GetRoleByPermissionName(r.Context(), permission.Permission)
}
