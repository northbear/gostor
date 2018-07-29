package gostor_test

import (
	"database/sql"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/northbear/gostor"
	_ "github.com/mattn/go-sqlite3"
)

const (
	DBNAME = "sqlite"
	DBPATH = "db/"
)

type filesupp struct{
	fname string
}

func (fs filesupp) GetFileName() string {
	if fs.fname == "" {
		fs.fname = path.Join(os.TempDir(), "testdb.sqlite3")
	}
	return fs.fname
}

type cnsupp struct{
	dbpath string
}

func (cns cnsupp) GetGostorContainer() gostor.Container {
	uri := fmt.Sprintf("file:%s?mode=rwc&cache=private", cns.dbpath)

	db, err := sql.Open("sqlite3", uri)
	if err != nil {
		fmt.Println(err)
	}
	
	container := gostor.New(db, "")
	return container
}

func TestGostor(t *testing.T) {
	tsupp := cnsupp{}
	gstor := tsupp.GetGostorContainer()
	
	if gstor == nil {
		t.Error("cannot create container")
	}
}

func TestNewToCreateContainer(t *testing.T) {
	var (
		db  *sql.DB
		err error
	)

	fs := filesupp{}
	uri := fmt.Sprintf("file:%s?mode=rwc&cache=private", fs.GetFileName())
	db, err = sql.Open("sqlite3", uri)
	if err != nil {
		t.Log(err)
	}
	
	resp := gostor.New(db, "")
	if _, ok := resp.(gostor.Container); !ok {
		t.Error("New do not return item with type of Container")
	} 	
}

type ProperStruct struct{
	Id int64
}

func TestRegisterToAcceptStuct(t *testing.T) {
	fname := path.Join(os.TempDir(), "testdb.sqlite3")
	defer os.Remove(fname)
	
	cntest := cnsupp{fname}
	cn := cntest.GetGostorContainer()
	
	err := cn.Register(ProperStruct{})
	if err != nil {
		t.Error("Cannot register proper struct: " + err.Error())
	}
}
