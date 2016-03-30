package v1

import (
	"github.com/jonaswouters/goclerk/models"
	"github.com/unrolled/render"
	"net/http"
)

func CreateOrganization(w http.ResponseWriter, r *http.Request) {

	render := render.New(render.Options{
		IndentJSON: true,
	})

	if r.Method == http.MethodPost {
		organization := &models.Organization{
			Name: "Test",
		}

		render.JSON(w, http.StatusOK, organization)
	} else {
		render.JSON(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
}
