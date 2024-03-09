package system

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	Conn *sql.DB
}

func NewStorage() (Storage, error) {
	conn, err := sql.Open("sqlite3", "./system/db.sqlite3?cache=shared&mode=rwc&_journal_mode=WAL&busy_timeout=10000")
	if err != nil {
		return Storage{}, err
	}
	return Storage{conn}, nil
}

func NewMemoryStorage() Storage {
	conn, err := sql.Open("sqlite3", "file::memory:?cache=shared&mode=rwc&_journal_mode=WAL&busy_timeout=10000")
	conn.SetMaxOpenConns(1)
	if err != nil {
        panic(err)
	}
	return Storage{conn}
}
