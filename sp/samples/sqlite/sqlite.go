package main

// #cgo LDFLAGS: -lsqlite3
// #include <stdio.h>
// #include <stdlib.h>
// #include <sqlite3.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func exec(db *C.sqlite3, statement string) error {
	sql := C.CString(statement)
	defer C.free(unsafe.Pointer(sql))
	rc := C.sqlite3_exec(db, sql, nil, nil, nil) // TODO callback
	if rc != C.SQLITE_OK {
		errorMessage := C.GoString(C.sqlite3_errmsg(db))
		return fmt.Errorf("sqlite3_exec returned code %d: %v", int(rc), errorMessage)
	}
	return nil
}

func main() {
	var db *C.sqlite3
	filename := C.CString("test.db")
	defer C.free(unsafe.Pointer(filename))
	rc := C.sqlite3_open(filename, &db)
	defer C.sqlite3_close(db)
	if rc != C.SQLITE_OK {
		panic(fmt.Sprintf("sqlite3_open returned code %d", int(rc)))
	}

	err := exec(db, "create table if not exists students (id integer not null primary key autoincrement, name text);")
	if err != nil {
		panic(err)
	}
	err = exec(db, "insert into students (name) values ('Johannes Weigend');")
	if err != nil {
		panic(err)
	}
	err = exec(db, "insert into students (name) values ('Bernhard Saumweber');")
	if err != nil {
		panic(err)
	}
	err = exec(db, "select * from students;")
	if err != nil {
		panic(err)
	}
}
