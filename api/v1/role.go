package v1

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (api *API) setupRoleRoutes(router *mux.Router) {
	role := router.PathPrefix("/admin").Subrouter()
	role.Use(api.authMiddleware.AdminOnly)

	// Role routes
	role.HandleFunc("/roles", api.roleHandler.GetRoles).Methods(http.MethodGet)
}
