package convert

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type TimeConvertor struct{}

var _ Converter = TimeConvertor{}
var _ FromValueConverter = TimeConvertor{}

func (tc TimeConvertor) ToValue(s string, vl reflect.Value) error {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return fmt.Errorf("wrong value to convert to time.Time: %s", s)
	}
	tm := time.Unix(v, 0)
	if vl.CanSet() {
		vl.Set(reflect.ValueOf(tm))
	} else {
		return fmt.Errorf("cannot set value. It is not addressable")
	}

	return nil
}

func (tc TimeConvertor) FromValue(rv reflect.Value) (string, error) {
	var tim time.Time
	var ok bool

	tim, ok = rv.Interface().(time.Time)
	if !ok {
		err := fmt.Errorf("wrong value type for conversion")
		return "", err
	}
	s := strconv.FormatInt(tim.Unix(), 10)

	return s, nil
}
