package main

import (
	"fmt"
	"reflect"
)

// START OMIT
func PrintVariableDetails(v interface{}) {
	typeof := reflect.TypeOf(v)
	fmt.Printf("The variable with type '%s' has the value '%v' and size %v\n", typeof.Name(), v, typeof.Size())
}

func main() {
	var someValue interface{}
	someValue = 2
	PrintVariableDetails(someValue)

	someValue = "abcd"
	PrintVariableDetails(someValue)

	if tmp, ok := someValue.(string); ok {
		fmt.Println("someValue is a string and has the value", tmp)
	}
}

// END OMIT
