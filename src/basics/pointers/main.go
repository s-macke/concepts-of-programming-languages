package main

import (
	"fmt"
	"unsafe"
)

// TODO
func main() {
	var x = 10
	var y = &x

	fmt.Println(x, *y)

	x = 20
	fmt.Println(x, *y)

	fmt.Printf("%p %p\n", &x, &y)

	z := []int{40, 50}

	y = &z[0]
	y = (*int)(unsafe.Add(unsafe.Pointer(y), 8))

	p := (*int)(unsafe.Pointer(uintptr(10) + 0))
	*p = 10

	//p := unsafe.Pointer(9)
	//reflect.N

	fmt.Println(p)

}
