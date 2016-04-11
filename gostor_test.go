package gostor_test

import (
	dbdrv "database/sql/driver"
	"testing"

	"github.com/northbear/gostor"
)

const (
	DBNAME = "sqlite"
	DBPATH = "db/"
)

var drv dbdrv.Driver

func GetTestContainer() gostor.Container {
	return gostor.New(drv, "")
}

func TestGostor(t *testing.T) {
	gstor := GetTestContainer()
	if gstor == nil {
		t.Error("cannot create container")
	}
}

