package main

import (
	"fmt"
	"unsafe"
)

/*
#cgo LDFLAGS: -lsqlite3
#include <stdio.h>
#include <stdlib.h>
#include <sqlite3.h>
int callback_cgo(void*, int, char**, char**);
*/
import "C"

var columnNamesPrinted bool

//export callback
func callback(_ unsafe.Pointer, columnCount C.int, columnTexts **C.char, columnNames **C.char) C.int {
	if !columnNamesPrinted {
		names := goSlice(columnCount, columnNames)
		separator := ""
		for i := 0; i < int(columnCount); i++ {
			fmt.Printf("| %-20s ", names[i])
			separator += fmt.Sprintf("|----------------------")
		}
		fmt.Printf("|\n")
		fmt.Printf("%s|\n", separator)
		columnNamesPrinted = true
	}
	texts := goSlice(columnCount, columnTexts)
	for i := 0; i < int(columnCount); i++ {
		fmt.Printf("| %-20s ", texts[i])
	}
	fmt.Printf("|\n")
	return C.SQLITE_OK
}

func exec(db *C.sqlite3, statement string) error {
	sql := C.CString(statement)
	defer C.free(unsafe.Pointer(sql))
	columnNamesPrinted = false
	rc := C.sqlite3_exec(db, sql, (*[0]byte)(C.callback_cgo), nil, nil)
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

func goSlice(argc C.int, argv **C.char) []string {
	length := int(argc)
	temp := (*[1 << 30]*C.char)(unsafe.Pointer(argv))[:length:length]
	slice := make([]string, length)
	for i, s := range temp {
		slice[i] = C.GoString(s)
	}
	return slice
}
