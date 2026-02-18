# Using PostgreSQL with Go net/http backend

Project structure:
- main.go - Main entrypoint. Runs api server and calls function creating DB connection
- database.go - DB connection logic
- handlers.go - REST api handlers
- services.go - logic for DB queries

Run a postgres database using docker

Install posgtres driver:
```bash
go get -u github.com/lib/pq
```

Import package and driver in database.sql:
```bash
import (
	"database/sql"
	_ "github.com/lib/pq"
)
```