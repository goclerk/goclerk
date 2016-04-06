[![License][License-Image]][License-URL] [![ReportCard][ReportCard-Image]][ReportCard-URL] [![Build][Build-Status-Image]][Build-Status-URL] [![Release][Release-Image]][Release-URL]
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

[License-URL]: https://github.com/jonaswouters/goclerk/blob/master/LICENSE
[License-Image]: https://img.shields.io/github/license/jonaswouters/goclerk.svg
[ReportCard-URL]: https://goreportcard.com/report/jonaswouters/goclerk
[ReportCard-Image]: https://goreportcard.com/badge/jonaswouters/goclerk
[Build-Status-URL]: https://travis-ci.org/jonaswouters/goclerk
[Build-Status-Image]: https://travis-ci.org/jonaswouters/goclerk.svg?branch=master
[Release-URL]: https://github.com/jonaswouters/goclerk/releases
[Release-Image]: https://img.shields.io/github/release/jonaswouters/goclerk.svg