package main

import "fmt"

type fnf func(fnf) fnf

func main() {
	ID := func(x fnf) fnf { return x }   // identity
	M := func(x fnf) fnf { return x(x) } // mocking bird

	fmt.Printf("ID = %p\n", ID)
	fmt.Printf("M(ID) = %p\n", M(ID))
	//fmt.Printf("M(M) = %p\n", M(M)) // omega
}
