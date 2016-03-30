package v1

import (
	"github.com/jonaswouters/goclerk/models"
	"net/http"
	"github.com/unrolled/render"
)

func CreateOrganization(w http.ResponseWriter, r *http.Request) {
	organization := &models.Organization{
		Name: "Test",
	}

	render := render.New(render.Options{
		IndentJSON: true,
	})

	render.JSON(w, http.StatusOK, organization)
}
