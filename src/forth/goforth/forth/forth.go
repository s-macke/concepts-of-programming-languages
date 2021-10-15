package forth

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var debug = false

// Forth contains the forth environment
type Forth struct {
	stack       stack
	returnStack stack // not used for the return addresses but for specific features like >r r> and loop parameters
	dict        dictionary
	heap        []int
	output      string
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
	f.dict = make(map[string]func())
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
