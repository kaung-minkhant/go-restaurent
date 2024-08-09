package router

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/kaung-minkhant/go-restaurent/auth"
	"github.com/kaung-minkhant/go-restaurent/database"
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
		fmt.Println("In auth middleware")
		token := r.Header.Get("x-access-token")
		if token == "" {
			AccessDeniedResponse(w)
			return
		}
		var err error
		refToken, err := getRefreshTokenFromContext(r)
		if err != nil {
			AccessDeniedResponse(w)
			return
		}

		var claims *auth.CustomClaims
		claims, err = auth.ValidateJWT(token)
		if err != nil {
			if strings.Contains(err.Error(), "token is expired") {
				if getRoutePattern(r) == "/auth/refresh" {
					fmt.Println("Refreshing")
					expiredClaims, err := auth.GetClaimsWithoutValidation(token)
					if err != nil {
						AccessDeniedResponse(w)
						return
					}
					fmt.Println("Expired claim", expiredClaims)
					claims = expiredClaims
				} else {
					fmt.Println("JWT validation failed", err)
					AccessDeniedResponse(w)
					return
				}
			} else {
				fmt.Println("JWT validation failed", err)
				AccessDeniedResponse(w)
				return
			}
		}
		// fmt.Println("Claim", claims)
		user, err := database.Db.GetUserByEmployeeId(r.Context(), claims.EmployeeID)
		if err != nil {
			SomethingWentWrongResponse(w)
			return
		}

		ctx := context.WithValue(r.Context(), ctxUserKey, &user)
		ctx = context.WithValue(ctx, ctxClaimsKey, claims)
		ctx = context.WithValue(ctx, ctxAccTokenKey, token)
		ctx = context.WithValue(ctx, ctxRefTokenKey, refToken)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RolePermissionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
		claims, err := getClaimsFromContext(r)
		if err != nil {
			AccessDeniedResponse(w)
			return
		}
		if rolePermission.Role != claims.EmployeeRole {
			AccessDeniedResponse(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}

var AuthWithRolePermission = chi.Chain(AuthMiddleware, RolePermissionMiddleware)
