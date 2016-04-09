package contact

import (
	"github.com/gorilla/schema"
	"github.com/jonaswouters/goclerk/models"
	"github.com/jonaswouters/goclerk/modules/setting"
	"github.com/jonaswouters/goclerk/modules/store"
	"github.com/jonaswouters/goclerk/routers/api"
	"github.com/siddontang/go/bson"
	"net/http"
)

// GetContacts get all contacts
func GetContacts(w http.ResponseWriter, r *http.Request) {

	var contacts []models.Contact
	err := store.GetDB().All(&contacts)

	if err != nil {
		setting.Renderer.JSON(w, http.StatusBadRequest, err)

		return
	}

	setting.Renderer.JSON(w, http.StatusOK, contacts)
}

// CreateContact create an contact
func CreateContact(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		setting.Renderer.JSON(w, http.StatusBadRequest, err)
	}

	contact := new(models.Contact)
	contact.ID = bson.NewObjectId()

	decoder := schema.NewDecoder()
	err = decoder.Decode(contact, r.PostForm)

	if err != nil {
		api.RenderError(w, http.StatusBadRequest, err)

		return
	}

	err = store.GetDB().Save(contact)

	if err != nil {
		api.RenderError(w, http.StatusBadRequest, err)

		return
	}

	setting.Renderer.JSON(w, http.StatusOK, contact)
}
