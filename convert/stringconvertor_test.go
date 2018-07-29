package convert

import (
	"reflect"
	"testing"
)

func TestStringConvertorToValue(t *testing.T) {
	var sample string = "sample string"
	var result string

	conv := StringConvertor{}
	err := conv.ToValue(sample, reflect.Indirect(reflect.ValueOf(&result)))
	if err != nil {
		t.Log("error in time of conversion:", err.Error())
	}

	if sample != result {
		t.Error("StringConvertor.ToValue does not convert properly")
	}
}
