package routers

import (
	"github.com/unrolled/render"
	"net/http"
)

// Home is the root endpoint for the web interface
func Home(w http.ResponseWriter, r *http.Request) {
	render := render.New(render.Options{
		IndentJSON: true,
	})

	render.JSON(w, http.StatusOK, map[string]string{"welcome": "This is rendered JSON!"})
}
