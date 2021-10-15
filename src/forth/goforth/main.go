package main

import (
	"github.com/s-macke/concepts-of-programming-languages/src/forth/goforth/forth"
)

func main() {
	forth.SetDebug(true)
	f := forth.NewForth()
	f.CmdLine()
}
