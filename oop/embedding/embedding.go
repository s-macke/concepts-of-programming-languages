package main

import (
	"fmt"
)

type Introducer interface{ Introduce() }
type Worker interface{ Work() }

type Person struct {
	Name string
}

// Person now implicitly satisfies Introducer
func (p Person) Introduce() {
	fmt.Printf("Hello, my name is %s\n", p.Name)
}

type Employee struct {
	Person
	Worker
	EmployeeID int
}

// EOE OMIT

// Employee will satisfy Worker even without this declaration
func (e Employee) Work() {
	fmt.Printf("<%s is working>\n", e.Name)
}

func main() {
	e := Employee{
		Person:     Person{"John"},
		EmployeeID: 1,
	}
	e.Introduce() // prints "Hello, my name is John"
	e.Work()      // does not require an implementation (but throws exception if not implemented)
}

// EOF OMIT
