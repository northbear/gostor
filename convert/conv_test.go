package convert

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"time"
)

type Person struct {
	Id      int64
	Name    string
	RegTime time.Time
}

var response []string = []string{"1", "Harry Urcen", "1081976400"} // 15.04.2004

func MoveValues(in []string, o interface{}) error {
	ov := reflect.Indirect(reflect.ValueOf(o))

	if ov.Type().Kind() != reflect.Struct {
		return fmt.Errorf("Output is not Struct")
	}

	for i := 0; i < ov.NumField(); i++ {
		fl := ov.Field(i)

		conv := stdConv.getConv(fl)
		if err := conv.ToValue(in[i], fl); err != nil {
			return fmt.Errorf(
				"wrong conversion value '%s' to field %d of struct %s",
				in[i], i, ov.Type().Name(),
			)
		}
	}
	return nil
}

func TestMoveValuesFromArrayToStruct(t *testing.T) {
	var p Person
	MoveValues(response, &p)

	isMoved := (strconv.FormatInt(p.Id, 10) == response[0] &&
		p.Name == response[1] &&
		strconv.FormatInt(p.RegTime.Unix(), 10) == response[2])

	if !isMoved {
		t.Errorf("data is not moved properly: %t", p)
	}
}
