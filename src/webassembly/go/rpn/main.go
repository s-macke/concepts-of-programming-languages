//go:build js
// +build js

package main

import (
	"fmt"
	"strconv"
	"strings"
	"syscall/js"
)

func Exec(input string) string {
	var s stack
	fields := strings.Fields(input)
	for _, expr := range fields {
		fmt.Println("Parse ", expr)
		switch expr {
		case "+":
			s.Push(s.Pop() + s.Pop())
		case "*":
			s.Push(s.Pop() * s.Pop())
		default:
			v, err := strconv.Atoi(expr)
			if err != nil {
				return "unknown expr."
			}
			s.Push(v)
		}
	}

	return strconv.Itoa(s.Pop())
}

func RPN(this js.Value, input []js.Value) interface{} {
	return js.ValueOf(Exec(input[0].String()))
}

func main() {
	// register function
	js.Global().Set("rpn", js.FuncOf(RPN))

	// prevent exit
	c := make(chan struct{}, 0)
	<-c
}
