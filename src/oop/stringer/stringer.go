package main

import "fmt"

type DoorOpen bool

func (d DoorOpen) String() string {
	if d == true {
		return "Door is open"
	} else {
		return "Door is closed"
	}
}

func main() {
	var d DoorOpen = false
	fmt.Println(d)
}
