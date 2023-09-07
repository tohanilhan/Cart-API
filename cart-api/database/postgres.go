package database

import (
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib" // load pgx driver for PostgreSQL
	"github.com/jmoiron/sqlx"

	"github.com/tohanilhan/Cart-API/pkg/repository"
	"github.com/tohanilhan/Cart-API/vars"
)

var Db *sqlx.DB

func PostgreSQLConnection() error {

	postgresConnURL, err := repository.ConnectionURLBuilder("sentinel-postgres")
	if err != nil {
		return err
	}

	Db, err = sqlx.Connect("pgx", postgresConnURL)
	if err != nil {
		return fmt.Errorf("error, not connected to database, %w", err)
	}

	Db.SetMaxOpenConns(vars.AppConfigs.PostgresqlMaxConnections)                           // the default is 0 (unlimited)
	Db.SetMaxIdleConns(vars.AppConfigs.PostgresqlMaxIdleConnections)                       // defaultMaxIdleConns = 2
	Db.SetConnMaxLifetime(time.Duration(vars.AppConfigs.PostgresqlMaxLifetimeConnections)) // 0, connections are reused forever

	if err := Db.Ping(); err != nil {
		defer Db.Close()
		return fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return nil

}
