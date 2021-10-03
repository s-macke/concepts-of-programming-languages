package main

import (
	"fmt"
	"os"
	"strconv"
)

var debug bool

type Forth struct {
	stack      Stack
	dictionary map[string]func()
	heap       []int
}

type ProgramState struct {
	forth              *Forth
	instructionPointer int
	funcs              []func()
}

func NewForthInterpreter() *Forth {
	var f = new(Forth)
	f.dictionary = make(map[string]func())
	f.FillDictionary()
	return f
}

func (f *Forth) booleanPush(result bool) {
	if result {
		f.stack.Push(1)
	} else {
		f.stack.Push(-1)
	}
}

func (f *Forth) FillDictionary() {
	f.dictionary["+"] = func() { f.stack.Push(f.stack.Pop() + f.stack.Pop()) }
	f.dictionary["*"] = func() { f.stack.Push(f.stack.Pop() * f.stack.Pop()) }
	f.dictionary["-"] = func() { temp := f.stack.Pop(); f.stack.Push(f.stack.Pop() - temp) }
	f.dictionary["/"] = func() { temp := f.stack.Pop(); f.stack.Push(f.stack.Pop() / temp) }

	f.dictionary[">"] = func() { temp := f.stack.Pop(); f.booleanPush(f.stack.Pop() > temp) }
	f.dictionary["<"] = func() { temp := f.stack.Pop(); f.booleanPush(f.stack.Pop() < temp) }
	f.dictionary["="] = func() { f.booleanPush(f.stack.Pop() == f.stack.Pop()) }
	f.dictionary["<>"] = func() { f.booleanPush(f.stack.Pop() == f.stack.Pop()) }

	f.dictionary["."] = func() { fmt.Println(f.stack.Pop()) }
	f.dictionary["cr"] = func() { fmt.Println() }
	f.dictionary["drop"] = func() { f.stack.Pop() }
	f.dictionary["dup"] = func() { temp := f.stack.Pop(); f.stack.Push(temp); f.stack.Push(temp) }
	f.dictionary["@"] = func() { f.stack.Push(f.heap[f.stack.Pop()]) }
	f.dictionary["!"] = func() { address := f.stack.Pop(); f.heap[address] = f.stack.Pop() }

	f.dictionary["words"] = func() {
		for k := range f.dictionary {
			fmt.Print(k, " ")
		}
		fmt.Println("")
	}

	f.dictionary[".s"] = func() {
		fmt.Printf("<%d> ", len(f.stack))
		for _, e := range f.stack {
			fmt.Print(e, " ")
		}
		fmt.Println("")
	}
}

func (f *Forth) GetWordFromDictionary(w string) func() {
	if f.dictionary[w] == nil {
		i, err := strconv.Atoi(w)
		if err != nil {
			fmt.Println("Error: Word '" + w + "' is neither a valid word nor a number")
			os.Exit(1)
		}
		return func() {
			f.stack.Push(i)
		}
	}
	return f.dictionary[w]
}

func (s *ProgramState) Execute() {
	if debug {
		fmt.Println("Execute")
	}
	s.instructionPointer = 0
	for {
		if s.instructionPointer >= len(s.funcs) {
			break
		}
		s.funcs[s.instructionPointer]()
		s.instructionPointer++
	}
}

func (f *Forth) Interpret(command string) {
	for _, w := range Lexer(command) {
		f.GetWordFromDictionary(w)()
	}
}

func (f *Forth) Exec(command string) {
	if debug {
		fmt.Printf("Exec '%s'\n", command)
	}
	Lexer(command).Parser().Compile(f).Execute()
}

func main() {
	debug = true
	f := NewForthInterpreter()

	//Exec("1 if 2 . then", f)
	//f.Exec("1 2 + .")
	//f.Exec("variable var1")
	//f.Exec("200 constant x")
	//f.Exec("0 if 2 . then 1 .")
	//f.Interpret("words")

	//f.Exec(": squared dup * ;")
	//f.Exec("3 squared .")

	f.Exec("( : rec 1 - dup . dup 0 > if recurse then ; )")
	f.Exec("20 rec")

	/*
		: fac2
		dup 0> if
		dup 1- recurse *
		else
		drop 1
		endif ;
		8 fac2 .
	*/

	/*
		f.Lexer("variable var1").Exec()
	*/
	//	f.Parse("2 var1 @ - .")
	//	f.ExecAll()

	//words := f.Lexer("2 3 + .")
	//words.Exec()
	//f.Exec("words")
	//f.Exec("1 2 3 .s")
}
