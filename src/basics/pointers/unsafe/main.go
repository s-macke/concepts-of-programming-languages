package main

import (
	"fmt"
	"unsafe"
)

func main() {
	z := [2]int64{10, 20} // fixed size array
	y := &z[0]            // pointer to first entry in array
	fmt.Println(y)
	y = (*int64)(unsafe.Add(unsafe.Pointer(y), 8)) // 64 Bit has 8 bytes
	fmt.Println(y)
	*y = 30 // changes second parameter in the array z
	fmt.Println(z)
}
