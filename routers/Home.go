package routers

import (
	"github.com/goclerk/goclerk/modules/middleware"
)

func Home(ctx *middleware.Context)  {
	ctx.Data["name"] = "Clerk"
	ctx.HTML(200, "hello")
}