package db

import (
	"fmt"
	"os"

	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/structs"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	Db     *sqlx.DB
	DbConf *structs.DbConfig
)

// Db function checks whether database connection is completed successfully or not. If not, panics.
func InitDb() {
	// load database config
	DbConf = &structs.DbConfig{
		Host:      os.Getenv("DB_HOST"),
		Port:      os.Getenv("DB_PORT"),
		Db:        os.Getenv("DB_DB"),
		Schema:    os.Getenv("DB_SCHEMA"),
		User:      os.Getenv("DB_USER"),
		Pass:      os.Getenv("DB_PASS"),
		TableName: os.Getenv("DB_TABLE_NAME"),
		SslMode:   "disable",
	}
	connStr := "postgres://%s:%s@%s:%s/%s?sslmode=%s"
	pq := fmt.Sprintf(connStr, DbConf.User, DbConf.Pass, DbConf.Host, DbConf.Port, DbConf.Db, DbConf.SslMode)
	Db = sqlx.MustConnect("postgres", pq)

	err := Db.Ping()

	if err != nil {
		panic(err)
	}

}
