package v1

import (
	"github.com/gorilla/mux"
	"github.com/jonaswouters/goclerk/routers/api/v1/organization"
	"net/http"
)

var (
	organizationPath = "/organization/"
)

// RegisterRoutes registers all v1 APIs routes to web application.
func RegisterRoutes(router *mux.Router) {
	// Organization endpoints
	router.HandleFunc(organizationPath, organization.GetOrganizations).Methods(http.MethodGet)
	router.HandleFunc(organizationPath, organization.CreateOrganization).Methods(http.MethodPost)
}
