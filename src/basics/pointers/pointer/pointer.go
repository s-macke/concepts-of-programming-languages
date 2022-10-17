package main

import "fmt"
import "reflect"

func main() {
    var x int32 = 10
    fmt.Println(reflect.TypeOf(x), reflect.TypeOf(&x))
    fmt.Println(reflect.TypeOf(x).Size(), reflect.TypeOf(&x).Size())
    fmt.Println(x, &x)  // shows value and memory address
    y := &x             // y is now a pointer to an int type. It points to the variable x
    *y = 20             // change of the variable x
    fmt.Println(x, &x)
}