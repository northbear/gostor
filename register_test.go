package gostor_test

import (
	"database/sql/driver"
	"testing"

	"github.com/northbear/gostor"
)

func TestRegisterImplication(t *testing.T) {
	var drv driver.Driver
	_ = gostor.New(drv, "")
}
