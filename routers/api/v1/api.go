package v1

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/jonaswouters/goclerk/routers/api/v1/organization"
)

var (
	organizationPath string = "/organization/"
)

// RegisterRoutes registers all v1 APIs routes to web application.
func RegisterRoutes(router *mux.Router) {
	// Organization endpoints
	router.HandleFunc(organizationPath, organization.GetOrganization).Methods(http.MethodGet)
	router.HandleFunc(organizationPath, organization.CreateOrganization).Methods(http.MethodPost)
}
