package datastore

import (
	"context"
	"database/sql"
	"errors"

	// MYSQL
	// _ "github.com/go-sql-driver/mysql"
	// PG

	// PG cgo
	// _ "github.com/jbarham/gopgsqldriver"
	// SQLITE
	_ "modernc.org/sqlite"
	// SQLITE cgo
	// _ "github.com/gwenn/gosqlite"
)

var db *sql.DB

func CreateConnection(ctx context.Context, uri string) (*sql.DB, error) {
	// _db, err := pgx.Connect(ctx, uri)
	_db, err := sql.Open("sqlite3", uri)
	if err != nil {
		return nil, err
	}

	return _db, nil
}

func GetConnection() (*sql.DB, error) {
	if db == nil {
		return nil, errors.New("there is no active connection")
	}

	return db, nil
}
