package api

import (
	"errors"
	"fmt"
	"github.com/jonaswouters/goclerk/modules/setting"
	"net/http"
)

// Errors
var (
	ErrAlreadyExists = errors.New("already exists")
	ErrNotFound      = errors.New("not found")
)

// RenderError renders errors for output
func RenderError(w http.ResponseWriter, status int, err error) {
	fmt.Println(err)
	setting.Renderer.JSON(w, status, map[string]string{"error": err.Error()})

	return
}
