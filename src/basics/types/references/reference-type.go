package main

import "fmt"

func main() {
	m1 := map[string]int{"x": 1, "y": 2}
	m2 := m1 // m2 points to the same map as m1
	m2["x"] = 100

	fmt.Println("m1:", m1) // map[x:100 y:2]
	fmt.Println("m2:", m2) // map[x:100 y:2]
}
