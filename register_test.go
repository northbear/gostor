package gostor

import (
	// "database/sql/driver"
	"bytes"
	"os"
	"os/exec"
	"path"
	"reflect"
	"strings"
	"testing"
)

type TestData struct{
	Id int64
}

type WrongData []string

type cliSqlite3 struct {
	dbpath string
}

func (tdb cliSqlite3) Exec(p string) {
	// buf := bytes.NewBuffer(make([]byte, 1024))
	cmd := exec.Command("sqlite3", tdb.dbpath, p)
	// cmd.Stdout, cmd.Stderr = buf, buf
	cmd.Run()
}

func (tdb cliSqlite3) Query(p string) string {
	buf := bytes.NewBuffer(make([]byte, 1024))

	cmd := exec.Command("sqlite3", tdb.dbpath, p)
	cmd.Stdout, cmd.Stderr = buf, buf

	cmd.Run()
	return buf.String() 
}

func TestTestSqlite3(t *testing.T) {
	dbfile := path.Join(os.TempDir(), "testdb.sql3")
	defer os.Remove(dbfile)

	testdb := cliSqlite3{dbfile}
	testdb.Exec("CREATE TABLE Users (Id INTEGER, Name STRING)")
	resp := testdb.Query(".schema")

	if !strings.Contains(resp, "TABLE Users") {
		t.Error("interacting with sqlite3 executable does not work")
	}
}

func TestRegisterReturnNilOnStructParam(t *testing.T) {
	cns := testGostor{}
	cn := cns.NewContainer()

	err := cn.Register(TestData{})
	if err != nil {
		t.Error("Register does not accept proper data")
	}
}

func TestRegisterReturnErrorOnNoneStructParam(t *testing.T) {
	cn := testGostor{}.NewContainer()

	err := cn.Register(WrongData{})
	if err != DataTypeError {
		t.Error("Register accepts improper data")
	}
}

func TestRegisterCreatesTableWithStructName(t *testing.T) {
	dbfile := path.Join(os.TempDir(), "testdb.sql3")
	defer os.Remove(dbfile)

	cli := cliSqlite3{dbfile}
	store := testGostor{dbfile}.NewContainer() 

	store.Register(TestData{})

	resp := cli.Query(".schema")

	if !strings.Contains(resp, "TABLE TestData") {
		t.Error("database table are not created properly")
	}
}

func TestFieldInfo(t *testing.T) {
	sinfo := reflect.TypeOf(TestStruct{})
	finfo := newFieldInfo(sinfo.Field(0))

	if finfo.NameType() != "Id INTEGER" {
		t.Error("incorrect NameType output") 
	}
}

func TestMetaInfo(t *testing.T) {
	minfo, err := newMetaInfo(TestStruct{})
	if err != nil {
		t.Error(err)
	}
	if minfo.FullNameTypeList() != "Id INTEGER, Name STRING" {
		t.Error("incorrent name-type list")
	} 
} 
