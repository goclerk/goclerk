# goclerk

A simple web app / api for small businesses and freelancers to help with all  the grunt work like document management, payment followup, time tracking, invoicing and more. 

Work in progress. No release yet

Trello board: https://trello.com/b/rcYSpU9g/goclerk

## Short term progress
- [x] Basic project
- [ ] Install & Setup command (create database/configuration)
  - [x] Create database
  - [x] Execute migrations
  - [x] Reset & Uninstall database
  - [x] Interactive commands 
  - [ ] Work with presetup database / users
  - [ ] Create users when working with super postgres user
  - [x] Generate config file after install
- [x] Basic models for Organizations, users, invoices and customers
- [x] Migrations
- [x] Create migration for basic models
- [ ] Decide on database tech (Postgres, BoltDB+Bleve, sqlite)
- [ ] Api basic functionality
  - [x] Form processing
  - [x] Form/Struct validation
  - [x] Json output
  - [ ] Authentication
 
## requirements

PostgreSQL

## Commands

### List of commands

Install the database
```bash
./goclerk setup install
```

Uninstall the database
```bash
./goclerk setup uninstall
```
Reset the database (for development or testing)
```bash
./goclerk setup reset
```

Run migrations
```bash
./goclerk migrate up
./goclerk migrate down
./goclerk migrate version
```

Run the app
```bash
./goclerk web
```
