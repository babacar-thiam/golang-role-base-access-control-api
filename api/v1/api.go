package v1

import (
	"github.com/gorilla/mux"
	"role-base-access-control-api/internal/auth"
	"role-base-access-control-api/internal/role"
	"role-base-access-control-api/internal/user"
)

type API struct {
	authHandler    *auth.Handler
	roleHandler    *role.Handler
	userHandler    *user.Handler
	authMiddleware *auth.Middleware
}

// NewAPI creates a new instance of API
func NewAPI(authHandler *auth.Handler, roleHandler *role.Handler, userHandler *user.Handler, authMiddleware *auth.Middleware) *API {
	return &API{authHandler: authHandler, roleHandler: roleHandler, userHandler: userHandler, authMiddleware: authMiddleware}
}

// SetupRoutes configures all the routes for v1 of the API
func (api *API) SetupRoutes(router *mux.Router) {
	// Add API version prefix
	v1 := router.PathPrefix("/api/v1").Subrouter()

	// Setup route groups
	api.setupAuthRoutes(v1)
	api.setupProtectedRoutes(v1)
}

// setupProtectedRoutes configures routes that require authentication
func (api *API) setupProtectedRoutes(router *mux.Router) {
	// Create protected sub-router
	protected := router.PathPrefix("").Subrouter()
	protected.Use(api.authMiddleware.AuthMiddleware)

	// Setup different route groups
	api.setupRoleRoutes(protected)
	api.setupUserRoutes(protected)
}
