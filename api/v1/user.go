package v1

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (api *API) setupUserRoutes(router *mux.Router) {
	usr := router.PathPrefix("/users").Subrouter()
	usr.Use(api.authMiddleware.AdminOnly)

	// User routes
	usr.HandleFunc("/all", api.userHandler.GetAllUsers).Methods(http.MethodGet)
	usr.HandleFunc("/{id}", api.userHandler.GetUser).Methods(http.MethodGet)

}
