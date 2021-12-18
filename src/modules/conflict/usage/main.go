package main

import (
	"fmt"
	"github.com/dep1"
	"github.com/dep2"
)

func main() {
	fmt.Println(dep1.MyVersion())
	fmt.Println(dep2.MyVersion())
}
