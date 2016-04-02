# gostor
Package for simple storing/restoring Go struct data in sqlite database


## Getting The Package

To get package to your local workspace execute command `go get github.com/northbear/gostor`

## Package Using

Import package to your application:

```Go
import (
    "github.com/northbear/gostor"
)
```

Define structure of data that you are going to store on database. The structure should
contain fields with given types: `integer`, `string`, `time.Time`, `bool`. Using another types and other structures is not allowed.

```Go
type TestStuct struct{
    Name   string
    Age    integer
    Date   time.Time
    Active bool
}
```


Initialize storage with type of database and path to database on the filesystem:

```Go
	gs := gostor.New("sqlite", "db/application.sqlite3")
```

Register structure supposed to be stored to certain storage:

```Go
err := gs.Register(TestStruct{})
	if err != nil {
	    t.Log("cannot register a structure in the container")
	}
```

Store data to storage

```Go
	sample := TestStruct{}
	err = gs.Store(sample)
	if err != nil {
 	    t.Log("cannot store a structure in the container")
	}
```

Retrieving data from storage

```Go
    var r interface{}
	r, err = gs.Restore(sample)
	if err != nil {
		t.Error("cannot restore structure: " + err.Error())
	}
	if _, ok := r.(TestStruct); !ok {
		t.Error("wrong type of structure restored")
	}
```
