// Package gostor allows to store data in Go structures in sqlite3 database
package gostor

import (
	"database/sql"
)

// Container interface represent an object to manipulate structs
type Container interface {
	// Register used to register structure going to be store in container
	Register(interface{}) error
	// Store previously registred struct
	Store(interface{}) error
	// Restore registred struct
	Restore(interface{}) (interface{}, error)
	// Restore set of registred structs
	RestoreSet(interface{}) ([]interface{}, error)
}

// New provides container for storing/restoring structures
func New(db *sql.DB, options string) Container {
	container := cntr{
		db:   db,
		options: options,
	}
	return container
}
