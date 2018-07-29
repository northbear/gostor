package convert

import (
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestTimeConvToValue(t *testing.T) {
	var sampleTime time.Time = time.Date(2004, 4, 15, 0, 0, 0, 0, time.UTC)
	var resultTime time.Time

	var conv TimeConvertor = TimeConvertor{}
	var sample string = strconv.FormatInt(sampleTime.Unix(), 10)

	err := conv.ToValue(sample, reflect.Indirect(reflect.ValueOf(&resultTime)))
	if err != nil {
		t.Log("error in time of conversion:", err.Error())
	}

	if !resultTime.Equal(sampleTime) {
		t.Error("TimeConvertor.ToValue returns wrong value")
	}
}

func TestTimeConvFromValue(t *testing.T) {
	var sampleTime time.Time = time.Date(2004, 4, 15, 0, 0, 0, 0, time.UTC)
	var sampleString string = strconv.FormatInt(sampleTime.Unix(), 10)

	var conv TimeConvertor = TimeConvertor{}

	resultString, err := conv.FromValue(reflect.Indirect(reflect.ValueOf(&sampleTime)))
	if err != nil {
		t.Log("error in time of conversion:", err.Error())
	}

	if !(resultString == sampleString) {
		t.Error("TimeConvertor.FromValue returns wrong value")
	}
}

func TestComplementarity(t *testing.T) {

}
