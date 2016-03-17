package cmd

import (
	"fmt"
	"os"

	migration "github.com/goclerk/goclerk/models/migrations"
	"gopkg.in/go-pg/migrations.v4"
	"github.com/codegangsta/cli"
	"gopkg.in/pg.v4"
)

var Install = cli.Command{
	Name:   "install",
	Usage:  "Install goclerk",
	Action: install,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "username, u",
			Usage: "Database username with database creation rights",

		},
		cli.StringFlag{
			Name:  "password, p",
			Usage: "Database password for username provided",
		},
	},
}

func init() {
	Migrations = append(Migrations, migration.CreateDatabase)
}

func install(ctx *cli.Context) {
	if (ctx.String("username") == "") {
		fmt.Fprintf(os.Stderr, "Username is required\n")
		os.Exit(1)
	}
	db := pg.Connect(&pg.Options{
		User:     ctx.String("username"),
		Password: ctx.String("password"),
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