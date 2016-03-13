package migrations

import (
	"fmt"
	"gopkg.in/go-pg/migrations.v4"
)

var CreateDatabase = migrations.Migration{
	Version: 1,
	Up: func(db migrations.DB) error {
		fmt.Println("creating database goclerk")
		_, err := db.Exec(`CREATE DATABASE goclerk`)
		return err
	},
	Down: func(db migrations.DB) error {
		fmt.Println("dropping datbase goclerk")
		_, err := db.Exec(`DROP DATABASE goclerk`)
		return err
	},
}
