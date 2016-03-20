package migrations

import (
	"fmt"
	"gopkg.in/go-pg/migrations.v4"
)

var Initialization = migrations.Migration{
	Version: 1,
	Up: func(db migrations.DB) error {
		fmt.Println("Upgrading Initialization migration")
		queries := []string{
			`CREATE TABLE public.organizations (
			 	id SERIAL NOT NULL,
				name CHARACTER VARYING(200) NOT NULL
			)`,
			`CREATE TABLE public.users (
				id SERIAL NOT NULL,
				username CHARACTER VARYING(200) NOT NULL,
				email CHARACTER VARYING(200) NOT NULL,
				password CHARACTER VARYING(200) NOT NULL
			)`,
		}
		for _, q := range queries {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		return nil
	},
	Down: func(db migrations.DB) error {
		fmt.Println("Downgrading Initialization migration")
		queries := []string{
			`DROP TABLE users`,
			`DROP TABLE organizations`,
		}
		for _, q := range queries {
			_, err := db.Exec(q)
			if err != nil {
				return err
			}
		}
		return nil
	},
}
