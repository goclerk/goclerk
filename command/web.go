package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"net/http"
	"strings"
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
	var address = strings.Join(
		[]string{
			"localhost",
			":",
			ctx.String("port"),
		},
		"",
	)
	http.HandleFunc("/", handler)
	fmt.Printf("Listening on http://%s", address)
	http.ListenAndServe(address, nil)

}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi there and welcome to GoClerk")
}
