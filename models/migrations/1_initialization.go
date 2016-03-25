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
			 	id SERIAL PRIMARY KEY,
				name TEXT NOT NULL
			)`,
			`CREATE TABLE public.users (
				id SERIAL PRIMARY KEY,
				username TEXT NOT NULL,
				email TEXT NOT NULL,
				password TEXT NOT NULL
			)`,
			`CREATE TABLE public.organization_users (
				organization_id INT REFERENCES public.organizations,
				user_id INT REFERENCES public.users
			)`,
			`CREATE TABLE public.customers (
				organization_id INT REFERENCES public.organizations,
				id SERIAL PRIMARY KEY,
				company_name TEXT NOT NULL,
				first_name TEXT NOT NULL,
				last_name TEXT NOT NULL,
				email TEXT NOT NULL,
				phone_number TEXT NOT NULL,
				vat_number TEXT NOT NULL
			)`,
			`CREATE TABLE public.customer_addresses (
				id SERIAL PRIMARY KEY,
				customer_id INT REFERENCES public.customers,
				address TEXT NOT NULL,
				postal_code TEXT NOT NULL,
				city TEXT NOT NULL,
				country TEXT NOT NULL
			)`,
			`CREATE TABLE public.invoices (
				organization_id INT REFERENCES public.organizations,
				id SERIAL PRIMARY KEY,
				number TEXT NOT NULL,
				customer_id INT REFERENCES public.customers,
				address TEXT NOT NULL,
				postal_code TEXT NOT NULL,
				city TEXT NOT NULL,
				country TEXT NOT NULL,
				amount INT NOT NULL,
				note TEXT NOT NULL,
				status TEXT NOT NULL,
				created_at TIMESTAMP NOT NULL,
				updated_at TIMESTAMP NOT NULL,
				invoice_date TIMESTAMP NOT NULL,
				due_date TIMESTAMP,
				paid_date TIMESTAMP
			)`,
			`CREATE TABLE public.invoice_details (
				invoice_id INT REFERENCES public.invoices,
				id SERIAL PRIMARY KEY,
				number TEXT NOT NULL,
				quantity INT NOT NULL,
				Description TEXT NOT NULL,
				amount INT NOT NULL,
				vat int NOT NULL
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
			`DROP TABLE invoice_details`,
			`DROP TABLE invoices`,
			`DROP TABLE customer_addresses`,
			`DROP TABLE customers`,
			`DROP TABLE organization_users`,
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
