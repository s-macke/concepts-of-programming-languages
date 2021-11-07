package main

import (
	"github.com/s-macke/concepts-of-programming-languages/src/forth/interpreter/forth"
)

func main() {
	forth.SetDebug(true)
	f := forth.NewForth()
	f.CmdLine()
}
