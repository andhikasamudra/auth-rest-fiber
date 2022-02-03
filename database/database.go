package database

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"os"
)

type Database struct {
	Conn *bun.DB
}

func New() *Database {
	dsn := os.Getenv("DB_URL")
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	return &Database{
		Conn: bun.NewDB(sqldb, pgdialect.New()),
	}
}
