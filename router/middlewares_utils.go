package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kaung-minkhant/go-restaurent/database"
	"github.com/kaung-minkhant/go-restaurent/database/models"
)

func getPermission(r *http.Request) (models.Permission, error) {
	routePattern := getRoutePattern(r)
	return database.Db.GetPermissionByMethodAndRoute(r.Context(), models.GetPermissionByMethodAndRouteParams{
		Method: r.Method,
		Route:  routePattern,
	})
}

func getRoutePattern(r *http.Request) string {
	return chi.RouteContext(r.Context()).RoutePattern()
}

func getRoleFromPermission(r *http.Request, permission models.Permission) (models.RolePermission, error) {
	return database.Db.GetRoleByPermissionName(r.Context(), permission.Permission)
}
