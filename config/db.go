package db

import (
	"database/sql"
	"fmt"
	"net/url"

	sqlserver "github.com/denisenkom/go-mssqldb"
	"github.com/jinzhu/gorm"
)

// OpenDB opens a DB connection with the mssql server.
func OpenDB() (*gorm.DB, error) {
	connString := url.URL{
		Scheme:   "sqlserver",
		Host:     "localhost:1433",
		User:     url.UserPassword("usernmae", "password"), // Please change me
		RawQuery: "database=example",
	}

	connector, err := sqlserver.NewConnector(connString.String())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	mssqlDB := sql.OpenDB(connector)

	db, err := gorm.Open("mssql", mssqlDB)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return db, nil
}
