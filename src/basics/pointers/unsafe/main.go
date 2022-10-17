package main

import (
	"fmt"
	"unsafe"
)

// TODO
func main() {
	//var x = 10
	//var y = &x

	//fmt.Println(x, *y)

	//x = 20
	//fmt.Println(x, *y)

	//fmt.Printf("%p %p\n", &x, &y)

	z := []int64{10, 20}
	y := &z[0]
	y = (*int64)(unsafe.Add(unsafe.Pointer(y), 8)) // 64 Bit it has 8 bytes
	*y = 30
	fmt.Println(z)

}
