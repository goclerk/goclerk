package cmd

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/jonaswouters/goclerk/models"
	"github.com/jonaswouters/goclerk/modules/setting"
	"github.com/jonaswouters/goclerk/modules/store"
	cliui "github.com/mitchellh/cli"
	"github.com/siddontang/go/bson"
)

// Setup command to install the configuration and reset the database
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

	database, _ := ui.Ask("Database filename (database.db):")

	if database != "" {
		setting.Settings.Database = database
	}

	// Create organization

	organizationName, _ := ui.Ask("Default organization name (default):")

	if organizationName == "" {
		organizationName = "default"
	}

	organization := new(models.Organization)
	organization.ID = bson.NewObjectId()
	organization.Name = organizationName
	err := store.GetDB().Save(organization)

	if err != nil {
		fmt.Printf("Failed to create organization %s: %s", organizationName, err.Error())

		return
	}

	// Create user

	username, _ := ui.Ask("Default username (admin):")

	if username == "" {
		username = "admin"
	}

	password, _ := ui.AskSecret("Default password (admin):")

	if password == "" {
		password = "admin"
	}

	user := new(models.User)
	user.ID = bson.NewObjectId()
	user.Username = username
	user.Password = password
	user.OrganizationIDs = []bson.ObjectId{organization.ID}
	err = store.GetDB().Save(user)

	if err != nil {
		fmt.Printf("Failed to create user %s: %s", username, err.Error())

		return
	}

	// Save settings
	setting.SaveSettings()
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
