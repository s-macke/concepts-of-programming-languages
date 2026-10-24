package main

import (
	"fmt"
	"reflect"
)

func main() {
	ADD := func(x, y int) int {
		return x + y
	}
	fmt.Println(ADD(1, 2)) // -> returns 3
	fmt.Println("The type of ADD is", reflect.TypeOf(ADD).Kind())
}
