package gostor

type Container interface {
	Register(interface{}) error
	Store(interface{}) error
	Restore(interface{}) (interface{}, error)
	RestoreSet(interface{}) ([]interface{}, error)
}

func New(db, path string) Container {
	return cntr{}
}

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
