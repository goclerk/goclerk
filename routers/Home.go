package routers

import (
	"github.com/goclerk/goclerk/modules/middleware"
)

func Home(ctx *middleware.Context) {
	ctx.Data["PageIsHome"] = true
	ctx.HTML(200, "hey")
}