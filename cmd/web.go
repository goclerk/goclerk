package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/codegangsta/negroni"

	"github.com/jonaswouters/goclerk/routers"
	apiv1 "github.com/jonaswouters/goclerk/routers/api/v1"

	"fmt"
	"github.com/go-ini/ini"
	"github.com/gorilla/mux"
	"github.com/jonaswouters/goclerk/modules/setting"
	"os"
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

// runWeb will serve the website and api
func runWeb(ctx *cli.Context) {
	// Settings
	cfg, err := ini.Load("settings.ini")

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}

	setting.Connection = setting.GetConnectionSettings(cfg.Section("database"))

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
