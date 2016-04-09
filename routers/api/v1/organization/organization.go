package organization

import (
	"fmt"
	"github.com/gorilla/schema"
	"github.com/jonaswouters/goclerk/models"
	"github.com/jonaswouters/goclerk/modules/store"
	"github.com/siddontang/go/bson"
	"github.com/unrolled/render"
	"net/http"
)

func GetOrganization(w http.ResponseWriter, r *http.Request) {

	render := render.New(render.Options{
		IndentJSON: true,
	})

	var organizations []models.Organization
	err := store.DB.All(&organizations)

	if err != nil {
		render.JSON(w, http.StatusBadRequest, err)
		fmt.Println(err)
		return
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
	organization.ID = bson.NewObjectId()

	decoder := schema.NewDecoder()
	err = decoder.Decode(organization, r.PostForm)

	if err != nil {
		render.JSON(w, http.StatusBadRequest, err)
		fmt.Println(err)
		return
	}

	err = store.DB.Save(organization)

	if err != nil {
		render.JSON(w, http.StatusBadRequest, err)
		fmt.Println(err)
		return
	}

	render.JSON(w, http.StatusOK, organization)
}
