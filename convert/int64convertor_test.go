package convert

import (
	"reflect"
	"strconv"
	"testing"
)

func TestInt64ConvertorToValue(t *testing.T) {
	var sample string = "5"
	var i int64

	var conv Int64Convertor = Int64Convertor{}

	err := conv.ToValue(sample, reflect.Indirect(reflect.ValueOf(&i)))
	if err != nil {
		t.Log("error in time of conversion:", err.Error())
	}

	if sample != strconv.FormatInt(i, 10) {
		t.Error("Int64Convertor.ToValue does not convert properly")
	}
}
