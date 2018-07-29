package gostor

import (
	"reflect"
)

type SqlTypeConterver interface{
	SqlType(reflect.Type) string
}

type SqlValueConterver interface{}
