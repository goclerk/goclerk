# goclerk

A simple web app / api for small businesses and freelancers to help with all  the grunt work like document management, payment followup, time tracking, invoicing and more. 

Work in progress. No release yet

## requirements

PostgreSQL

## Commands

### List of commands

Install the database (Will be interactive later on)
```bash
./goclerk setup install -u <username> -p <password>
```

Uninstall the database
```bash
./goclerk setup uninstall -u <username> -p <password>
```
Reset the database (for development or testing)
```bash
./goclerk setup reset -u <username> -p <password>
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