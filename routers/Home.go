package routers

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GoClerk!\n"))
}
