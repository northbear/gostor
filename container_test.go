package gostor

import (
	"database/sql"
	"fmt"
	"testing"

)


type TestStruct struct{
	Id int64
	Name string
}

var (
	_ Container = cntr{}
	drv *sql.DB
)

type testGostor struct{
	dbpath string
}

func (cns testGostor) NewContainer() Container {
	uri := fmt.Sprintf("file:%s?mode=rwc&cache=private", cns.dbpath)

	db, err := sql.Open("sqlite3", uri)
	if err != nil {
		fmt.Println(err)
	}
	
	container := New(db, "")
	return container
}


func TestRegister(t *testing.T) {
	cns := testGostor{}
	gs := cns.NewContainer()

	err := gs.Register(TestStruct{})

	if err != nil {
		t.Error("cannot register structure in container")
	}
}

func TestStore(t *testing.T) {
	cns := testGostor{}
	gs := cns.NewContainer()
	err := gs.Register(TestStruct{})
	if err != nil {
		t.Log("cannot register a structure in the container")
	}

	sample := TestStruct{}
	err = gs.Store(sample)
	if err != nil {
		t.Error("cannot store a structure in the container")
	}
}

func TestRestore(t *testing.T) {
	cns := testGostor{}
	gs := cns.NewContainer()
	err := gs.Register(TestStruct{})
	if err != nil {
		t.Log("cannot register a structure in the container")
	}
	sample := TestStruct{}
	err = gs.Store(sample)
	if err != nil {
		t.Log("cannot store a structure in the container")
	}

	var r interface{}
	r, err = gs.Restore(sample)
	if err != nil {
		t.Error("cannot restore structure: " + err.Error())
	}
	if _, ok := r.(TestStruct); !ok {
		t.Error("wrong type of structure restored")
	}
}

func TestRestoreSet(t *testing.T) {
	cns := testGostor{}
	gs := cns.NewContainer()
	err := gs.Register(TestStruct{})
	if err != nil {
		t.Log("cannot register a structure in the container")
	}
	sample := TestStruct{}
	err = gs.Store(sample)
	if err != nil {
		t.Log("cannot store a structure in the container")
	}

	var r []interface{}

	r, err = gs.RestoreSet(sample)
	if err != nil {
		t.Error("cannot restore structure: " + err.Error())
	}
	if _, ok := r[0].(TestStruct); !ok {
		t.Error("wrong type of structure restored")
	}
}
