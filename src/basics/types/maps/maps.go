package main

import "fmt"

func main() {
	m := make(map[string]string) // Initialize an empty map
	m["foo"] = "bar"             // insert a key-value pair into map

	value, ok := m["asd"]  // check if key is present, return parameters can be ignored with "_"
	fmt.Println(ok, value) // returns false, value is just an empty string ""

	value, ok = m["foo"] // returns true, value contains "bar"
	fmt.Println(ok, value)

	for k, v := range m {
		fmt.Println(k, v)
	}
}
