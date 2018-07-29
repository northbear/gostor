package gostor

import (
	"errors"
	"fmt"
	"reflect"
)

var DataTypeError error = errors.New("wrong type of data for register")

type metainfo struct {
	Name   string
	fields []fieldinfo
}

type fieldinfo struct {
	reflect.StructField
}

func newFieldInfo(fs reflect.StructField) fieldinfo {
	fi := fieldinfo{fs}
	return fi
}

func (fi fieldinfo) NameType() string {
	return fi.Name + " " + fi.TypeConverted()
}

func (fi fieldinfo) TypeConverted() string {
	var sqlType string
	switch fi.Type.Kind() {
	case reflect.String:
		sqlType = "STRING"
	case reflect.Int, reflect.Int64:
		sqlType = "INTEGER"
	}

	return sqlType
}

func (mi *metainfo) FullNameTypeList() string {
	list := ""
	for i, fi := range mi.fields {
		if i != 0 {
			list = list + ", "
		}
		list = list + fi.NameType()
	}
	return list
}

func newMetaInfo(data interface{}) (metainfo, error) {
	mi := metainfo{}

	tpe := reflect.TypeOf(data)
	if tpe.Kind() != reflect.Struct {
		return mi, DataTypeError
	}
	mi.Name = tpe.Name()
	mi.fields = make([]fieldinfo, tpe.NumField())
	for i := range mi.fields {
		mi.fields[i] = newFieldInfo(tpe.Field(i))
	}

	return mi, nil
}

type SqlCmd struct{}

func (cmd SqlCmd) CreateTable(tname, flist string) string {
	return fmt.Sprintf("CREATE TABLE %s (%s)", tname, flist)
}

func (c cntr) Register(data interface{}) error {
	mi, err := newMetaInfo(data)
	if err != nil {
		return err
	}
	c.registered[mi.Name] = mi
	cmd := SqlCmd{}
	_, err = c.db.Exec(cmd.CreateTable(mi.Name, mi.FullNameTypeList()))
	return err
}
