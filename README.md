[![License][License-Image]][License-URL] [![ReportCard][ReportCard-Image]][ReportCard-URL] [![Build][Build-Status-Image]][Build-Status-URL] [![Release][Release-Image]][Release-URL]
# goclerk

A simple web app / api for small businesses and freelancers to help with all  the grunt work like document management, payment followup, time tracking, invoicing and more. 

Work in progress. No release yet

Trello board: https://trello.com/b/rcYSpU9g/goclerk

## Short term progress
- [x] Basic project
- [x] Install & Setup command (create configuration)
  - [x] Reset & Uninstall database
  - [x] Interactive commands 
  - [x] Create a default user and organization
  - [x] Generate config file after install
- [x] Basic models for Organizations, users, invoices and customers
- [x] Decide on database tech -> BoltDB+Bleve
- [ ] Basic web interface 
  - [ ] Authentication
- [ ] Api basic functionality
  - [x] Form processing
  - [x] Form/Struct validation
  - [x] Json output
  - [ ] Authentication
  - [ ] post & validate json instead of form


## Commands

### List of commands

Install the database
```bash
./goclerk setup install
```

Reset the database (for development or testing)
```bash
./goclerk setup reset
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
