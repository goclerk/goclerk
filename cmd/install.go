package cmd

import (
	"fmt"
	"os"

	"github.com/goclerk/goclerk/models/migrations"
	"github.com/codegangsta/cli"
	"gopkg.in/pg.v4"
	"github.com/goclerk/goclerk/modules/setting"
)

var Install = cli.Command{
	Name:   "install",
	Usage:  "Install goclerk",
	Action: install,
}

func init() {
	migrations.Register(migrations.CreateDatabase)
}

func install(ctx *cli.Context) {
	db := pg.Connect(&pg.Options{
		User:     setting.Connection.Username,
		//Database: setting.Connection.Database,
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