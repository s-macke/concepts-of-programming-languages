package main

/*
#cgo LDFLAGS: -lsqlite3
#include <stdlib.h>
#include <sqlite3.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func exec(db *C.sqlite3, sql string) {
	csql := C.CString(sql)
	defer C.free(unsafe.Pointer(csql))
	rv := C.sqlite3_exec(db, csql, nil, nil, nil)
	if rv != C.SQLITE_OK {
		fmt.Printf("error: %s\n", C.GoString(C.sqlite3_errmsg(db)))
	}
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
	exec(db, "create table if not exists students (id integer not null primary key autoincrement, name text);")
	exec(db, "insert into students (name) values ('Your Name');")
	exec(db, "insert into students (name) values ('Another Name');")
}
