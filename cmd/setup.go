package cmd

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/jonaswouters/goclerk/modules/setting"
	cliui "github.com/mitchellh/cli"
	"gopkg.in/go-pg/migrations.v4"
	"gopkg.in/pg.v4"
)

var (
	ignorePasswordFlag = cli.StringFlag{
		Name:  "ignore-password, p",
		Usage: "Ignore password prompt",
	}
	usernameFlag = cli.StringFlag{
		Name:  "username, u",
		Usage: "Database username with database creation rights",
	}
)

var Setup = cli.Command{
	Name:  "setup",
	Usage: "Goclerk setup tools",
	Subcommands: []cli.Command{
		{
			Name:   "install",
			Usage:  "Install goclerk",
			Action: install,
			Flags: []cli.Flag{
				usernameFlag,
				ignorePasswordFlag,
			},
		},
		{
			Name:   "reset",
			Usage:  "Reset the database",
			Action: reset,
			Flags: []cli.Flag{
				usernameFlag,
				ignorePasswordFlag,
			},
		},
		{
			Name:   "uninstall",
			Usage:  "uninstall goclerk",
			Action: uninstall,
			Flags: []cli.Flag{
				usernameFlag,
				ignorePasswordFlag,
			},
		},
	},
}

// install will create the database and run all migrations
func install(ctx *cli.Context) {
	username, password := getUsernameAndPassword(ctx)

	db := pg.Connect(&pg.Options{
		User:     username,
		Password: password,
	})

	ui := &cliui.BasicUi{Writer: os.Stdout, Reader: os.Stdin}

	database, _ := ui.Ask("Database name:")

	var err error

	_, err = db.Exec(fmt.Sprintf(`CREATE DATABASE %s`, database))

	if err != nil {
		db.Close()
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	} else {
		fmt.Printf("Database %s created", database)

		db = pg.Connect(&pg.Options{
			User:     username,
			Password: password,
			Database: database,
		})

		_, _, err = migrations.RunMigrations(db, []migrations.Migration{}, "init")

		db.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error()+"\n")
			os.Exit(1)
		}
	}
}

// reset will drop the database schema and run all migrations again
func reset(ctx *cli.Context) {
	setting.LoadSettings()
	username, password := getUsernameAndPassword(ctx)

	db := pg.Connect(&pg.Options{
		User:     username,
		Password: password,
		Database: setting.Connection.Database,
	})

	queries := []string{
		`DROP SCHEMA public cascade`,
		`CREATE SCHEMA public`,
	}
	for _, q := range queries {
		_, err := db.Exec(q)
		if err != nil {
			db.Close()
			fmt.Fprintf(os.Stderr, err.Error()+"\n")
			os.Exit(1)
		}
	}

	_, _, err := migrations.RunMigrations(db, []migrations.Migration{}, "init")

	db.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}

	fmt.Printf("Database %s reset", setting.Connection.Database)
}

// uninstall will drop the database and remove configuration
func uninstall(ctx *cli.Context) {
	setting.LoadSettings()
	username, password := getUsernameAndPassword(ctx)

	db := pg.Connect(&pg.Options{
		User:     username,
		Password: password,
	})

	var err error
	_, err = db.Exec(fmt.Sprintf(`DROP DATABASE %s`, setting.Connection.Database))

	db.Close()
	
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	} else {
		fmt.Printf("Database %s deleted", setting.Connection.Database)
	}
}

// getUsernameAndPassword get username and password via flags or user input
func getUsernameAndPassword(ctx *cli.Context) (username string, password string) {
	ui := &cliui.BasicUi{Writer: os.Stdout, Reader: os.Stdin}

	username = ctx.String("username")

	if ctx.String("username") == "" {
		username, _ = ui.Ask("What is the username of PostgreSQL with create database permissions?")
	}

	if ctx.Bool("ignore-password") == false {
		password, _ = ui.AskSecret(fmt.Sprintf("Password for user %s?", username))
	}

	return username, password
}
