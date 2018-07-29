package gostor

import (
	"fmt"
	"log"
	"reflect"
)

type Value reflect.Value
type Type reflect.Kind

func (t Type) String() string {
	return reflect.Kind(t).String()
}

type SqlType string
type SqlValue string

type ValueConverter interface {
	ToValue(s string) Value
	FromValue(v Value) string
}

type TypeConverter interface {
	ToType(s string) Type
	FromType(v Type) string
}

type StringConvertor struct{}

func getType(k reflect.Kind) reflect.Type {
	var ret reflect.Type
	switch k {
	case reflect.String:
		ret = reflect.TypeOf("")
	case reflect.Int64:
		ret = reflect.TypeOf(int64(0))
	case reflect.Int:
		ret = reflect.TypeOf(int(0))
	default:
		log.Panic("unconverted type: " + k.String())
	}
	return ret
}

func getNewValue(k reflect.Kind) interface{} {
	var v interface{}

	switch k {
	case reflect.String:
		v = ""
	case reflect.Int64:
		v = int64(0)
	case reflect.Int:
		v = int(0)
	default:
		log.Panic("unconverted type: " + k.String())
	}
	return v
}

func (sc StringConvertor) ToValue(s string) Value {
	var result Value
	return result
}

func (sc StringConvertor) FromValue(v Value) string {
	result := ""
	return result
}

func (sc StringConvertor) ToType(s string) Type {
	result := Type(reflect.String)
	return result
}

func (sc StringConvertor) FromType(v Type) string {
	result := "STRING"
	return result
}

type StandardConvertor struct {
	typ reflect.Kind
}

type convInfo struct {
	goType  reflect.Kind
	sqlType string
	goFmt   string
	sqlFmt  string
}

var convInfoSet map[reflect.Kind]convInfo = map[reflect.Kind]convInfo{
	reflect.Int:    {reflect.Int, "INTEGER", "%d", "%d"},
	reflect.Int64:  {reflect.Int64, "INTEGER(8)", "%d", "%d"},
	reflect.String: {reflect.String, "TEXT", "%s", "%s"},
}

func NewStandardConvertor(rt reflect.Kind) StandardConvertor {
	sc := StandardConvertor{rt}
	return sc
}

func (cv StandardConvertor) ToType(s string) Type {
	result := Type(cv.typ)
	return result
}

func (cv StandardConvertor) FromType(v Type) string {
	var result string
	if v == Type(cv.typ) {
		result = convInfoSet[cv.typ].sqlType
	}
	return result
}

func (cv StandardConvertor) ToValue(s string) Value {
	// var (
	// 	v Value
	// )
	ci := convInfoSet[cv.typ]

	v := getNewValue(ci.goType)
	fmt.Sscanf(s, ci.goFmt, &v)
	if test, ok := v.(int64); ok {
		print(s)
		print(test)
	}
	return Value(reflect.ValueOf(v))
}

func (cv StandardConvertor) FromValue(v Value) string {
	var format string
	format = convInfoSet[cv.typ].sqlFmt
	return fmt.Sprintf(format, v)
}
