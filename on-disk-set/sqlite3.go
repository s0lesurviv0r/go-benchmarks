package set

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// SQLite3 uses the DB to back a set
type SQLite3 struct {
	path string
	db   *sql.DB
}

func NewSQLite3(path string) Set {
	return &SQLite3{
		path: path,
	}
}

func (s *SQLite3) Open() error {
	db, err := sql.Open("sqlite3", s.path)
	if err != nil {
		return err
	}
	s.db = db

	// Create table if it doesn't exist
	_, err = s.db.Exec(`CREATE TABLE IF NOT EXISTS set_data (
		id INTEGER PRIMARY KEY
	)`)
	if err != nil {
		return err
	}

	return nil
}

func (s *SQLite3) Close() error {
	return s.db.Close()
}

func (s *SQLite3) Add(v uint64) error {
	_, err := s.db.Exec("INSERT OR IGNORE INTO set_data (id) VALUES (?)", v)
	return err
}

func (s *SQLite3) Delete(v uint64) error {
	_, err := s.db.Exec("DELETE FROM set_data WHERE id = ?", v)
	return err
}

func (s *SQLite3) Exists(v uint64) (bool, error) {
	var count int
	err := s.db.QueryRow("SELECT COUNT(*) FROM set_data WHERE id = ?", v).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error querying existence: %w", err)
	}
	return count > 0, nil
}
