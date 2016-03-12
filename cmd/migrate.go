package cmd

import (
	"github.com/codegangsta/cli"
)

var Migrate = cli.Command{
	Name:   "migrate",
	Usage:  "Install & upgrade database to latest version",
	Subcommands: []cli.Command{
		{
			Name:  "version",
			Usage: "version of the database",
			Action: getversion,
		},
		{
			Name:  "up",
			Usage: "upgrade the database",
			Action: upgrade,
		},
		{
			Name:  "down",
			Usage: "downgrade the database",
			Action: downgrade,
		},
	},
}

func upgrade(ctx *cli.Context) {
	println("Will ugprade the datbase in the future")
}

func downgrade(ctx *cli.Context) {
	println("Will downgrade the datbase in the future")
}

func getversion(ctx *cli.Context) {
	println("Still version 0")
}