package cmd

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"gopkg.in/pg.v4"
	"gopkg.in/go-pg/migrations.v4"
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
		cli.BoolFlag{
			Name:  "uninstall",
			Usage: "Delete the database",
		},
	},
}

func install(ctx *cli.Context) {
	if ctx.String("username") == "" {
		fmt.Fprintf(os.Stderr, "Username is required\n")
		os.Exit(1)
	}
	db := pg.Connect(&pg.Options{
		User:     ctx.String("username"),
		Password: ctx.String("password"),
	})

	var err error
	if ctx.Bool("uninstall") {
		_, err = db.Exec(`DROP DATABASE goclerk`)

		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error()+"\n")
			os.Exit(1)
		} else {
			fmt.Printf("Database goclerk deleted")
		}
	} else {
		_, err = db.Exec(`CREATE DATABASE goclerk`)

		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error()+"\n")
			os.Exit(1)
		} else {
			fmt.Printf("Database goclerk created")
		}
	}
}
