package cmd

import (
	"database/sql"
	"fmt"
	"net/url"

	sqlserver "github.com/denisenkom/go-mssqldb"
	"github.com/jinzhu/gorm"
)

// SQLServerDBConfig database configuration.
type SQLServerDBConfig struct {
	SQLServerDB   string `envconfig:"SQLSERVER_DB"   default:"example"` // SQLServer Database name
	SQLServerHost string `envconfig:"SQLSERVER_HOST" default:"0.0.0.0"` // SQLServer Host
	SQLServerPort string `envconfig:"SQLSERVER_PORT" default:"3306"`    // SQLServer Port
	SQLServerUser string `envconfig:"SQLSERVER_USER" default:"root"`    // SQLServer User
}

// SQLServerPassConfig database password.
type SQLServerPassConfig struct {
	SQLServerPass string `envconfig:"SQLSERVER_PASS" default:"YourStrong@Passw0rd"` // SQLServer Pass
}

// User struct
type User struct {
	id   int64
	name string
	age  int
}

// OpenDB opens a DB connection with the mssql server.
func (cfg SQLServerDBConfig) OpenDB(passwordCfg SQLServerPassConfig, env string) (*gorm.DB, error) {
	connString := url.URL{
		Scheme:   "sqlserver",
		Host:     cfg.SQLServerHost + ":" + cfg.SQLServerPort,
		User:     url.UserPassword(cfg.SQLServerUser, passwordCfg.SQLServerPass),
		RawQuery: "database=" + cfg.SQLServerDB,
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

func main() {
	fmt.Println("in here")
}
