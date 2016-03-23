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
				name TEXT NOT NULL
			)`,
			`CREATE TABLE public.users (
				id SERIAL NOT NULL,
				username TEXT NOT NULL,
				email TEXT NOT NULL,
				password TEXT NOT NULL
			)`,
			`CREATE TABLE public.customers (
				id SERIAL NOT NULL,
				company_name TEXT NOT NULL,
				first_name TEXT NOT NULL,
				last_name TEXT NOT NULL,
				email TEXT NOT NULL,
				phone_number TEXT NOT NULL,
				vat_number TEXT NOT NULL
			)`,
			`CREATE TABLE public.addresses (
				id SERIAL NOT NULL,
				address TEXT NOT NULL,
				postal_code TEXT NOT NULL,
				city TEXT NOT NULL,
				country TEXT NOT NULL,
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
			`DROP TABLE customers`,
			`DROP TABLE addresses`,
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
