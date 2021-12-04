//go:build js
// +build js

package main

import (
	"syscall/js"
)

func add(this js.Value, input []js.Value) interface{} {
	result := input[0].Float() + input[1].Float()
	return js.ValueOf(result)
}

func main() {
	document := js.Global().Get("document")
	p := document.Call("createElement", "p")
	p.Set("innerHTML", "Hello WASM from Go!")
	document.Get("body").Call("appendChild", p)

	// register function
	js.Global().Set("add", js.FuncOf(add))

	// prevent exit
	c := make(chan struct{}, 0)
	<-c
}
