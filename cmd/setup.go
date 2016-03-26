package cmd

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"gopkg.in/pg.v4"
	"gopkg.in/go-pg/migrations.v4"
)

var Setup = cli.Command{
	Name:   "setup",
	Usage:  "Goclerk setup tools",
	Subcommands: []cli.Command{
		{
			Name:  "install",
			Usage: "Install goclerk",
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
		},
		{
			Name:  "reset",
			Usage: "Reset the database",
			Action: reset,
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
		},
		{
			Name:  "uninstall",
			Usage: "uninstall goclerk",
			Action: uninstall,
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
		},
	},
}

// install will create the database and run all migrations
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

	_, err = db.Exec(`CREATE DATABASE goclerk`)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	} else {
		fmt.Printf("Database goclerk created")

		db = pg.Connect(&pg.Options{
			User:     ctx.String("username"),
			Password: ctx.String("password"),
			Database: "goclerk",
		})

		_, _, err = migrations.RunMigrations(db, []migrations.Migration{}, "init")

		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error() + "\n")
			os.Exit(1)
		}
	}

}

// reset will drop the database schema and run all migrations again
func reset(ctx *cli.Context) {
	if ctx.String("username") == "" {
		fmt.Fprintf(os.Stderr, "Username is required\n")
		os.Exit(1)
	}
	db := pg.Connect(&pg.Options{
		User:     ctx.String("username"),
		Password: ctx.String("password"),
		Database: `goclerk`,
	})

	queries := []string{
		`DROP SCHEMA public cascade`,
		`CREATE SCHEMA public`,
	}
	for _, q := range queries {
		_, err := db.Exec(q)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error()+"\n")
			os.Exit(1)
		}
	}

	_, _, err := migrations.RunMigrations(db, []migrations.Migration{}, "init")

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error() + "\n")
		os.Exit(1)
	}

	fmt.Printf("Database goclerk reset")


}

// uninstall will drop the database and remove configuration
func uninstall(ctx *cli.Context) {
	if ctx.String("username") == "" {
		fmt.Fprintf(os.Stderr, "Username is required\n")
		os.Exit(1)
	}
	db := pg.Connect(&pg.Options{
		User:     ctx.String("username"),
		Password: ctx.String("password"),
	})

	var err error
	_, err = db.Exec(`DROP DATABASE goclerk`)

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	} else {
		fmt.Printf("Database goclerk deleted")
	}

}
