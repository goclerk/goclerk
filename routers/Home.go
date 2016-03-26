package routers

import (
	"github.com/jonaswouters/goclerk/modules/middleware"
)

func Home(ctx *middleware.Context) {
	ctx.Data["name"] = "Clerk"
	ctx.HTML(200, "hello")
}
