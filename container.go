package gostor

import ()

type cntr struct{}

func (c cntr) Register(interface{}) error {
	return nil
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

