package convert

import (
	"reflect"
	"testing"
	"time"
)

func TestProvidingInt64Convertor(t *testing.T) {
	var i int64 = 365

	conv := stdConv.getConv(reflect.ValueOf(i))
	if _, ok := conv.(Int64Convertor); !ok {
		t.Error("StandardConvertor does not provide proper convertor for Int64")
	}
}

func TestProvidingStringConvertor(t *testing.T) {
	var i string = "test string"

	conv := stdConv.getConv(reflect.ValueOf(i))
	if _, ok := conv.(StringConvertor); !ok {
		t.Error("StandardConvertor does not provide proper convertor for String")
	}
}

func TestProvidingTimeConvertor(t *testing.T) {
	var i time.Time = time.Now()

	conv := stdConv.getConv(reflect.ValueOf(i))
	if _, ok := conv.(TimeConvertor); !ok {
		t.Error("StandardConvertor does not provide proper convertor for time.Time")
	}
}
