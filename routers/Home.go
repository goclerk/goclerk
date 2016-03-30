package routers

import (
	"net/http"
	"github.com/unrolled/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render := render.New(render.Options{
		IndentJSON: true,
	})

	render.JSON(w, http.StatusOK, map[string]string{"welcome": "This is rendered JSON!"})
}
