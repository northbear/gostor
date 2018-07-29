package convert

import (
	"fmt"
	"reflect"
)

type Converter interface {
	ToValue(string, reflect.Value) error
}

type FromValueConverter interface {
	FromValue(reflect.Value) (string, error)
}

type StructConvSet map[string]Converter
type BasicConvSet map[reflect.Kind]Converter

type ConvertorKit struct {
	forBasic  BasicConvSet
	forStruct StructConvSet
}

type DefaultConvertor struct{}

var _ Converter = DefaultConvertor{}

func (dc DefaultConvertor) ToValue(s string, vl reflect.Value) error {
	return fmt.Errorf("no proper convertor for field type %s", vl.Type().Kind())
}

func (dc DefaultConvertor) FromValue(vl reflect.Value) (string, error) {
	return "", fmt.Errorf("no proper convertor for field type %s", vl.Type().Kind())
}

func (sc ConvertorKit) getConv(vl reflect.Value) Converter {
	var conv Converter

	kind := vl.Type().Kind()

	switch kind {
	case reflect.Int64, reflect.String:
		conv = sc.forBasic[kind]
	case reflect.Struct:
		structName := vl.Type().PkgPath() + "." + vl.Type().Name()
		conv = sc.forStruct[structName]
	}

	if conv == nil {
		return DefaultConvertor{}
	}

	return conv
}

var stdConv ConvertorKit = ConvertorKit{
	BasicConvSet{
		reflect.Int64:  Int64Convertor{},
		reflect.String: StringConvertor{},
	},
	StructConvSet{"time.Time": TimeConvertor{}},
}
