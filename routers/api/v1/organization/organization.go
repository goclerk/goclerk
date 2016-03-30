package organization

import (
	"github.com/jonaswouters/goclerk/models"
	"github.com/unrolled/render"
	"net/http"
)

func GetOrganization(w http.ResponseWriter, r *http.Request) {

	render := render.New(render.Options{
		IndentJSON: true,
	})

	organization := &models.Organization{
		Name: "GetTest",
	}

	render.JSON(w, http.StatusOK, organization)
}

func CreateOrganization(w http.ResponseWriter, r *http.Request) {

	render := render.New(render.Options{
		IndentJSON: true,
	})

	organization := &models.Organization{
		Name: "PostTest",
	}

	render.JSON(w, http.StatusOK, organization)
}

