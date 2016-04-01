package gostor

import (
	"testing"
)

const (
	DBNAME = "sqlite"
	DBPATH = "db/"
)

type TestStruct struct{}

func TestGostor(t *testing.T) {
	var gstor Container = New(DBNAME, DBPATH)
	if gstor == nil {
		t.Error("cannot create container")
	}
}

func TestRegister(t *testing.T) {
	gs := New(DBNAME, DBPATH)

	err := gs.Register(TestStruct{})

	if err != nil {
		t.Error("cannot register structure in container")
	}
}

func TestStore(t *testing.T) {
	gs := New(DBNAME, DBPATH)
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
	gs := New(DBNAME, DBPATH)
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
	gs := New(DBNAME, DBPATH)
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
