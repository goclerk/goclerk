package cmd

import (
	"github.com/urfave/cli"
	"github.com/urfave/negroni"

	"github.com/jonaswouters/goclerk/routers"
	apiv1 "github.com/jonaswouters/goclerk/routers/api/v1"

	"github.com/gorilla/mux"
	"github.com/jonaswouters/goclerk/modules/setting"
)

// Web is the web interface and api command
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

// runWeb will serve the website and api
func runWeb(ctx *cli.Context) {
	// Load settings
	setting.LoadSettings()

	n := newNegroni()

	router := mux.NewRouter()
	router.HandleFunc("/", routers.Home)
	api := router.PathPrefix("/api").Subrouter()

	apiv1.RegisterRoutes(api)

	n.UseHandler(router)
	n.Run(":4000")
}

// newNegroni initializes Negroni instance.
func newNegroni() *negroni.Negroni {
	n := negroni.New()

	return n
}
