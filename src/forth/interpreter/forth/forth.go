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
	returnStack stack // not used for the return addresses but for specific features such as >r r> and loop parameters
	dict        dictionary
	heap        []int
}

// SetDebug sets the debug level of the Forth lexer, parser and compiler
func SetDebug(value bool) {
	debug = value
}

// NewForth creates the Forth environment
func NewForth() *Forth {
	var f = new(Forth)
	f.dict = NewDictionary()
	return f
}

// NewRunContext creates a new run context
func (f *Forth) NewRunContext(words Parser) runContext {
	return runContext{
		forth:              f,
		instructionPointer: 0,
		words:              words,
		output:             "",
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

// Exec Interprets every word in the command line.
func (f *Forth) Exec(command string) (result string, err error) {
	parser := NewParser(Lexer(command))
	parser.Parse()
	c := f.NewRunContext(parser)
	defer func() {
		if r := recover(); r != nil {
			result = c.output
			err = errorHandler(r)
		}
	}()
	c.Run()
	return c.output, nil
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
