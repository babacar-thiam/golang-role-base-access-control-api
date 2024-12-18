package v1

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (api *API) setupAuthRoutes(router *mux.Router) {
	auth := router.PathPrefix("/auth").Subrouter()

	auth.HandleFunc("/register", api.authHandler.Register).Methods(http.MethodPost)
	auth.HandleFunc("/login", api.authHandler.Login).Methods(http.MethodPost)
}
