package organization

import (
	"github.com/gorilla/schema"
	"github.com/jonaswouters/goclerk/models"
	"github.com/jonaswouters/goclerk/modules/setting"
	"github.com/jonaswouters/goclerk/modules/store"
	"github.com/jonaswouters/goclerk/routers/api"
	"github.com/siddontang/go/bson"
	"net/http"
)

// GetOrganizations get all organizations
func GetOrganizations(w http.ResponseWriter, r *http.Request) {

	var organizations []models.Organization
	err := store.GetDB().All(&organizations)

	for i := 0; i < len(organizations); i++ {
		organization := &organizations[i]

		var organizationUsers []models.OrganizationUsers
		err = store.GetDB().Find("OrganizationId", organization.ID, &organizationUsers)

		for i := 0; i < len(organizationUsers); i++ {
			organizationUser := &organizationUsers[i]
			var user models.User
			err = store.GetDB().One("ID", organizationUser.UserID, &user)

			organization.Users = append(organization.Users, user)
		}
	}

	if err != nil {
		api.RenderError(w, http.StatusBadRequest, err)

		return
	}

	setting.Renderer.JSON(w, http.StatusOK, organizations)
}

// CreateOrganization create an organization
func CreateOrganization(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		setting.Renderer.JSON(w, http.StatusBadRequest, err)
	}

	organization := new(models.Organization)
	organization.ID = bson.NewObjectId()

	decoder := schema.NewDecoder()
	err = decoder.Decode(organization, r.PostForm)

	if err != nil {
		api.RenderError(w, http.StatusBadRequest, err)

		return
	}

	err = store.GetDB().Save(organization)

	if err != nil {
		api.RenderError(w, http.StatusBadRequest, err)

		return
	}

	setting.Renderer.JSON(w, http.StatusOK, organization)
}
