package system

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	// _ "github.com/mattn/go-sqlite3"
	"github.com/tursodatabase/go-libsql"
)

type Storage struct {
	Conn *sql.DB
}

func NewStorage() (Storage, error, func()) {
	dbName := "local.db"
	dir, err := os.MkdirTemp("", "libsql-*")
	if err != nil {
		return Storage{}, fmt.Errorf("Error creating temp dir: %w", err), nil
	}
	defer os.RemoveAll(dir)

	dbPath := filepath.Join(dir, dbName)

	connector, err := libsql.NewEmbeddedReplicaConnector(
		dbPath,
		TURSO_URL,
		libsql.WithAuthToken(TURSO_TOKEN),
	)
	if err != nil {
		return Storage{}, fmt.Errorf("Error creating connector: %w", err), nil
	}
	// defer connector.Close()

	db := sql.OpenDB(connector)
	// defer db.Close()
	// conn, err := sql.Open("libsql", TURSO_URL)
	// conn, err := sql.Open("sqlite3", "file::memory:?cache=shared&mode=rwc&_journal_mode=WAL&busy_timeout=10000")

	var clean func() = func() {
		os.RemoveAll(dir)
		connector.Close()
		db.Close()
	}

	return Storage{db}, nil, clean
}

func NewMemoryStorage() Storage {
	conn, err := sql.Open("libsql", "file::memory:?cache=shared&mode=rwc&_journal_mode=WAL&busy_timeout=10000")
	conn.SetMaxOpenConns(1)
	if err != nil {
		panic(err)
	}
	return Storage{conn}
}
