package gostor

import (
	"database/sql"
)

type cntr struct{
	db *sql.DB
	options string
	registered map[string]metainfo
}

func (c cntr) Store(interface{}) error {
	return nil
}

func (c cntr) Restore(d interface{}) (interface{}, error) {
	return d, nil
}

func (c cntr) RestoreSet(d interface{}) ([]interface{}, error) {
	set := []interface{}{ d }
	return set, nil
}

