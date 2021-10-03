package forth

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var debug = false

// Forth contains the forth environment
type Forth struct {
	stack      stack
	dictionary map[string]func()
	heap       []int
	output     string
}

type programState struct {
	name               string
	forth              *Forth
	instructionPointer int
	funcs              []func()
}

// SetDebug sets the debug level of the Forth lexer, parser and compiler
func SetDebug(value bool) {
	debug = value
}

// NewForth creates the Forth environment
func NewForth() *Forth {
	var f = new(Forth)
	f.dictionary = make(map[string]func())
	f.fillDictionary()
	return f
}

func (f *Forth) booleanPush(result bool) {
	if result {
		f.stack.Push(-1)
	} else {
		f.stack.Push(0)
	}
}

func (f *Forth) fillDictionary() {
	f.dictionary["+"] = func() { f.stack.Push(f.stack.Pop() + f.stack.Pop()) }
	f.dictionary["*"] = func() { f.stack.Push(f.stack.Pop() * f.stack.Pop()) }
	f.dictionary["-"] = func() { temp := f.stack.Pop(); f.stack.Push(f.stack.Pop() - temp) }
	f.dictionary["/"] = func() { temp := f.stack.Pop(); f.stack.Push(f.stack.Pop() / temp) }

	f.dictionary[">"] = func() { temp := f.stack.Pop(); f.booleanPush(f.stack.Pop() > temp) }
	f.dictionary["<"] = func() { temp := f.stack.Pop(); f.booleanPush(f.stack.Pop() < temp) }
	f.dictionary["="] = func() { f.booleanPush(f.stack.Pop() == f.stack.Pop()) }
	f.dictionary["<>"] = func() { f.booleanPush(f.stack.Pop() == f.stack.Pop()) }

	f.dictionary["."] = func() { f.output += fmt.Sprintf("%d ", f.stack.Pop()) }
	f.dictionary["cr"] = func() { f.output += fmt.Sprintln() }
	f.dictionary["drop"] = func() { f.stack.Pop() }
	f.dictionary["dup"] = func() { temp := f.stack.Pop(); f.stack.Push(temp); f.stack.Push(temp) }
	f.dictionary["@"] = func() { f.stack.Push(f.heap[f.stack.Pop()]) }
	f.dictionary["!"] = func() { address := f.stack.Pop(); f.heap[address] = f.stack.Pop() }
	f.dictionary["bye"] = func() { panic("Program stopped") }

	f.dictionary["words"] = func() {
		for k := range f.dictionary {
			f.output += fmt.Sprint(k, " ")
		}
		f.output += fmt.Sprintln("")
	}

	f.dictionary[".s"] = func() {
		f.output += fmt.Sprintf("<%d> ", len(f.stack))
		for _, e := range f.stack {
			f.output += fmt.Sprint(e, " ")
		}
	}
}

func (f *Forth) getWordFromDictionary(w string) func() {
	if f.dictionary[w] == nil {
		i, err := strconv.Atoi(w)
		if err != nil {
			panic("Word '" + w + "' is neither a valid word nor a number")
		}
		return func() {
			f.stack.Push(i)
		}
	}
	return f.dictionary[w]
}

func (s *programState) Execute() {
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

func errorHandler(r interface{}) error {
	switch r.(type) {
	case string:
		return errors.New(r.(string))
	case error:
		return r.(error)
	default:
		return errors.New("Unknown error type")
	}
}

// Interpret just runs every word in the command line. No compiling
func (f *Forth) Interpret(command string) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			result = f.output
			err = errorHandler(r)
		}
	}()
	f.output = ""
	for _, w := range lexer(command) {
		f.getWordFromDictionary(w)()
	}
	return f.output, nil
}

// Exec compiles the given command and runs it
func (f *Forth) Exec(command string) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			result = f.output
			err = errorHandler(r)
		}
	}()

	if debug {
		fmt.Printf("Exec '%s'\n", command)
	}

	f.output = ""
	lexer(command).Parser().compile(f).Execute()
	return f.output, nil
}

func (f *Forth) CmdLine() {
	fmt.Println("Forth Command Line Interpreter")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		result, err := f.Exec(scanner.Text())
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(result)
		}
	}
}
