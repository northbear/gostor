package gostor

import (
	"reflect"
	"testing"
)

func TestConverters(t *testing.T) {
	t.Error("Contervers required")
}

type TypeTestData struct {
	goType  reflect.Kind
	sqlType string
	tconv   TypeConverter
}

var typesTable []TypeTestData = []TypeTestData{
	{reflect.String, "TEXT", NewStandardConvertor(reflect.String)},
	{reflect.Int, "INTEGER", NewStandardConvertor(reflect.Int)},
	{reflect.Int64, "INTEGER(8)", NewStandardConvertor(reflect.Int64)},
}

func TestConversionFromType(t *testing.T) {
	var td TypeTestData
	for _, td = range typesTable {
		actual := td.tconv.FromType(Type(td.goType))
		if actual != td.sqlType {
			name := reflect.TypeOf(td.tconv).Name()
			t.Errorf("wrong %s conversion from type %s to %s. Actual result: %s", name, td.goType.String(), td.sqlType, actual)
		}
	}
}

func TestConversionToType(t *testing.T) {
	for _, td := range typesTable {
		actual := td.tconv.ToType(td.sqlType)
		if actual != Type(td.goType) {
			name := reflect.TypeOf(td.tconv).Name()
			t.Errorf("wrong %s conversion to type %s from %s. Actual result: %s", name, td.goType.String(), td.sqlType, reflect.Kind(actual).String())
		}
	}
}

type ValueTestData struct {
	goValue  reflect.Value
	sqlValue string
	vconv    ValueConverter
}

var valuesTable []ValueTestData = []ValueTestData{
	{reflect.ValueOf("Hello!"), "Hello!", NewStandardConvertor(reflect.String)},
	{reflect.ValueOf(5), "5", NewStandardConvertor(reflect.Int)},
}

func TestConversionFromValue(t *testing.T) {
	t.Error("write test for values conversion")
}

// func TestConversionToValue(t *testing.T) {
// 	// var actual Value
// 	for _, vt := range valuesTable {
// 		actual := reflect.Value(vt.vconv.ToValue(vt.sqlValue))
// 		if vt.goValue.Interface() != actual.Interface() {
// 			tn := vt.goValue.Type().Name()
// 			t.Errorf("wrong value conversion for type %s. Expected %s, Actual %s", tn, vt.sqlValue, actual.String())
// 		}
// 	}
// }

func TestConversionToValueInt(t *testing.T) {
	value := int64(5)
	conv := NewStandardConvertor(reflect.Int64)
	actual := reflect.Value(conv.ToValue("5"))

	if value != actual.Int() {
		tn := "Int"
		t.Errorf("wrong value conversion for type %s. Expected %s, Actual %d", tn, value, actual.Int())
	}
}

func TestValueConverters(t *testing.T) {

	var _ ValueConverter = StringConvertor{}
	var _ TypeConverter = StringConvertor{}

	// t.Error("wrong value conversion with string convertor")
}
