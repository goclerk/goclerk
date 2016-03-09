package command

import (
	"github.com/codegangsta/cli"
	"gopkg.in/macaron.v1"

	"github.com/goclerk/goclerk/routers"


)

var Web = cli.Command{
	Name:   "web",
	Usage:  "Start GoClerk web server",
	Action: runWeb,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "config/config.ini",
			Usage: "Configuration file path",
		},
		cli.StringFlag{
			Name:  "port, p",
			Value: "4000",
			Usage: "Port number",
		},
	},
}

func runWeb(ctx *cli.Context) {
	m := newMacaron()

	m.Get("/", routers.Home)

	m.Run()

}

// newMacaron initializes Macaron instance.
func newMacaron() *macaron.Macaron {
	m := macaron.New()
	m.Use(macaron.Renderer())

	return m
}
