package v1

import (
	"github.com/gorilla/mux"
)

// RegisterRoutes registers all v1 APIs routes to web application.
func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/organization/", CreateOrganization)
}
