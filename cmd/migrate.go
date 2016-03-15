package cmd

import (
	"fmt"
	"os"

	"github.com/goclerk/goclerk/models/migrations"
	"github.com/codegangsta/cli"
	"gopkg.in/pg.v4"
	"github.com/goclerk/goclerk/modules/setting"
)

var Migrate = cli.Command{
	Name:   "migrate",
	Usage:  "Install & upgrade database to latest version",
	Subcommands: []cli.Command{
		{
			Name:  "version",
			Usage: "version of the database",
			Action: getVersion,
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

func init() {
	//migrations.Register(migrations.CreateDatabase)
}

func upgrade(ctx *cli.Context) {
	db := pg.Connect(&pg.Options{
		User:     setting.Connection.Username,
		Database: setting.Connection.Database,
	})

	oldVersion, newVersion, err := migrations.Run(db, "up")
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}

	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
}

func downgrade(ctx *cli.Context) {
	println("Will downgrade the datbase in the future")
}

func getVersion(ctx *cli.Context) {
	println("Still version 0")
}