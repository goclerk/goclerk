package cmd

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/jonaswouters/goclerk/modules/setting"
	cliui "github.com/mitchellh/cli"
)

var Setup = cli.Command{
	Name:  "setup",
	Usage: "Goclerk setup tools",
	Subcommands: []cli.Command{
		{
			Name:   "install",
			Usage:  "Install goclerk",
			Action: install,
		},
		{
			Name:   "reset",
			Usage:  "Reset the database",
			Action: reset,
		},
	},
}

// install will create the database and run all migrations
func install(ctx *cli.Context) {
	setting.LoadSettings()
	ui := &cliui.BasicUi{Writer: os.Stdout, Reader: os.Stdin}

	database, _ := ui.Ask("Database filename:")

	if database != "" {
		// Save settings
		setting.Settings.Database = database
		setting.SaveSettings()
	}
}

// reset will drop the database schema and run all migrations again
func reset(ctx *cli.Context) {
	setting.LoadSettings()

	err := os.Remove(setting.Settings.Database)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Database %s reset", setting.Settings.Database)
}
