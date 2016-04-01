package organization

import (
	"github.com/gorilla/schema"
	"github.com/jonaswouters/goclerk/models"
	"github.com/unrolled/render"
	"gopkg.in/pg.v4"
	"net/http"
	"github.com/jonaswouters/goclerk/modules/setting"
)

func GetOrganization(w http.ResponseWriter, r *http.Request) {

	render := render.New(render.Options{
		IndentJSON: true,
	})

	db := pg.Connect(&pg.Options{
		User:     setting.Connection.Username,
		Database: setting.Connection.Database,
	})

	var organizations []models.Organization
	err := db.Model(&organizations).Select()

	if err != nil {
		render.JSON(w, http.StatusBadRequest, err)
	}

	render.JSON(w, http.StatusOK, organizations)
}

func CreateOrganization(w http.ResponseWriter, r *http.Request) {

	render := render.New(render.Options{
		IndentJSON: true,
	})

	err := r.ParseForm()

	if err != nil {
		render.JSON(w, http.StatusBadRequest, err)
	}

	organization := new(models.Organization)

	decoder := schema.NewDecoder()
	err = decoder.Decode(organization, r.PostForm)

	if err != nil {
		render.JSON(w, http.StatusBadRequest, err)
	}

	db := pg.Connect(&pg.Options{
		User:     setting.Connection.Username,
		Database: setting.Connection.Database,
	})

	err = db.Create(organization)
	if err != nil {
		render.JSON(w, http.StatusBadRequest, err)
	}

	render.JSON(w, http.StatusOK, organization)
}
