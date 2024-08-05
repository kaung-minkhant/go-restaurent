package router

import (
	"net/http"
	"time"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func InitRouter() *chi.Mux {
	r := chi.NewRouter()
	setupMiddleWares(r)

	setupHealthCheck(r)

	r.Route("/menu_items", setupMenuItemRoutes)

	r.Route("/categories", setupCategoryRoutes)

	r.Route("/sub_categories", setupSubCategoryRoutes)

	r.Route("/auth", setupAuthRoutes)

	r.Route("/roles", setupRoleRoutes)

	r.Route("/test", setupTestRoutes)

	return r
}

func setupTestRoutes(r chi.Router) {
	r.Get("/{id}", makeHandlerFunc(handleGetTest))
}

func setupMenuItemRoutes(r chi.Router) {
	r.Get("/", makeHandlerFunc(handleGetAllMenuItems))
	r.Get("/{menuItemId}", makeHandlerFunc(handleGetMenuItemById))
	r.With(AuthMiddleware).Post("/", makeHandlerFunc(handleCreateMenuItem))
	r.With(AuthMiddleware).Delete("/{menuItemId}", makeHandlerFunc(handleDeleteMenuItem))
	r.With(AuthMiddleware).Patch("/{menuItemId}", makeHandlerFunc(handleUpdateMenuItem))
}

func setupCategoryRoutes(r chi.Router) {
	r.Get("/", makeHandlerFunc(handleGetAllCategory))
	r.Post("/", makeHandlerFunc(handleCreateCategory))
	r.Delete("/{categoryId}", makeHandlerFunc(handleDeleteCategory))
	r.Patch("/{categoryId}", makeHandlerFunc(handleUpdateCategory))
}

func setupSubCategoryRoutes(r chi.Router) {
	r.Get("/", makeHandlerFunc(handleGetAllSubCategory))
	r.Post("/", makeHandlerFunc(handleCreateSubCategory))
	r.Delete("/{subCategoryId}", makeHandlerFunc(handleDeleteSubCategory))
	r.Patch("/{subCategoryId}", makeHandlerFunc(handleUpdateSubCategory))
}

func setupRoleRoutes(r chi.Router) {
	r.With(AuthMiddleware).With(ValidateRolePermissionMiddleware).Post("/", makeHandlerFunc(handleCreateRole))
}

func setupAuthRoutes(r chi.Router) {
	r.Post("/signin", makeHandlerFunc(handleSignIn))
	r.With(AuthMiddleware).Get("/refresh", makeHandlerFunc(handleRefreshToken))
	r.With(AuthMiddleware).With(ValidateRolePermissionMiddleware).Post("/signup", makeHandlerFunc(handleSignUp))
}

func setupMiddleWares(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// timeout
	r.Use(middleware.Timeout(10 * time.Second))
}

func setupHealthCheck(r *chi.Mux) {
	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("All ok"))
	})
}
