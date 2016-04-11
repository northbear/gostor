package gostor

import (
	"database/sql/driver"
)

type Container interface {
	Register(interface{}) error
	Store(interface{}) error
	Restore(interface{}) (interface{}, error)
	RestoreSet(interface{}) ([]interface{}, error)
}

func New(db driver.Driver, path string) Container {
	return cntr{}
}
