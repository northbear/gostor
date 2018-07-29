package convert

import "reflect"

type StringConvertor struct{}

var _ Converter = StringConvertor{}

func (sc StringConvertor) ToValue(s string, vl reflect.Value) error {
	vl.SetString(s)

	return nil
}

func (sc StringConvertor) FromValue(rv reflect.Value, s string) error {
	return nil
}
