### About

I use this program to test queries using gorm's mssql dialect.

### Dependencies

This assumes you have [SQL Server running in a Docker container](https://docs.microsoft.com/en-us/sql/linux/quickstart-install-connect-docker?view=sql-server-ver15&pivots=cs1-bash).

### Update DB user config

Be sure to update the username and password in `config/db.go`:

```go
User:     url.UserPassword("username", "password"), // Please change me
```

### Create DB

To create and migrate the db, the `db/example.sql` file provides the necessary commands. These can be run in Azure Data Studio or in a sql-server console.

### Build

```sh
cd go-mssql-example
make build
```

### Run

```sh
bin/example
```

### License

MIT
