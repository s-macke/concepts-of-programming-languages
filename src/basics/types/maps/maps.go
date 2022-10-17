package main

import "fmt"

func main() {
	m := map[string]string{"foo": "bar"}
	_, ok := m["asd"]
	fmt.Println(ok)
	_, ok = m["foo"]
	fmt.Println(ok)
	for k, v := range m {
		fmt.Println(k, v)
	}
}
