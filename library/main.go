package main

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -lmylib
// #include "mylib.h"
// #include <stdlib.h>
import "C"
import "unsafe"

func main() {
	s := C.CString("Hello Lib!")
	defer C.free(unsafe.Pointer(s))
	C.myprint(s)
}
