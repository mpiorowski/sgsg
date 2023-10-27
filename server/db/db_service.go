package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func Connect() error {
	var err error
	Db, err = sql.Open("sqlite3", "./db/db.sqlite3?cache=shared&mode=rwc&_journal_mode=WAL&busy_timeout=10000")
	// Db.SetMaxOpenConns(1)
	if err != nil {
		return err
	}
	return nil
}

func ConnectTest() error {
	var err error
    Db, err = sql.Open("sqlite3", "file::memory:?cache=shared&mode=rwc&_journal_mode=WAL&busy_timeout=10000")
    Db.SetMaxOpenConns(1)
	if err != nil {
		return err
	}
	return nil
}
