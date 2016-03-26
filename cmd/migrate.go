package cmd

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	migration "github.com/jonaswouters/goclerk/models/migrations"
	"github.com/jonaswouters/goclerk/modules/setting"
	"gopkg.in/go-pg/migrations.v4"
	"gopkg.in/pg.v4"
)

var (
	Migrations []migrations.Migration
)

var Migrate = cli.Command{
	Name:  "migrate",
	Usage: "Install & upgrade database to latest version",
	Subcommands: []cli.Command{
		{
			Name:   "version",
			Usage:  "version of the database",
			Action: getVersion,
		},
		{
			Name:   "up",
			Usage:  "upgrade the database",
			Action: upgrade,
		},
		{
			Name:   "down",
			Usage:  "downgrade the database",
			Action: downgrade,
		},
	},
}

func init() {
	Migrations = append(Migrations, migration.Initialization)
}

// upgrade will migrate the database to the latest version
func upgrade(ctx *cli.Context) {
	db := pg.Connect(&pg.Options{
		User:     setting.Connection.Username,
		Database: setting.Connection.Database,
	})

	oldVersion, newVersion, err := migrations.RunMigrations(db, Migrations, "up")
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

// downgrade will migrate the database back one version
func downgrade(ctx *cli.Context) {
	db := pg.Connect(&pg.Options{
		User:     setting.Connection.Username,
		Database: setting.Connection.Database,
	})

	oldVersion, newVersion, err := migrations.RunMigrations(db, Migrations, "down")
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
	}

	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
}

// getVersion prints the current version
func getVersion(ctx *cli.Context) {
	println("Still version 0")
}
