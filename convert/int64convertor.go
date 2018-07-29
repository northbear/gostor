package convert

import (
	"fmt"
	"reflect"
	"strconv"
)

type Int64Convertor struct{}

var _ Converter = Int64Convertor{}

func (ic Int64Convertor) ToValue(s string, vl reflect.Value) error {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return fmt.Errorf("wrong value to convert to int64: %s", s)
	}
	vl.SetInt(v)

	return nil
}

func (ic Int64Convertor) FromValue(rv reflect.Value, s string) error {
	return nil
}
