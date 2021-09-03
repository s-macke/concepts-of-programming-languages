package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Stack []int

type Forth struct {
	stack      Stack
	dictionary map[string]func(w *Words)
	heap       []int
}

type Words struct {
	forth       *Forth
	parsedIndex int
	parsedwords []string
}

func (f *Forth) NewParser() Words {
	return Words{
		parsedIndex: 0,
		parsedwords: make([]string, 0),
		forth:       f,
	}
}

func ExtractFuncFromParser(words *Words) Words {
	neww := words.forth.NewParser()
	for {
		w := words.GetNextParsedWord()
		if w == nil {
			panic(w)
		}
		if *w == ";" {
			return neww
		}
		neww.parsedwords = append(neww.parsedwords, *w)
	}
}

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str int) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack.
func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic("Stack is empty")
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element
	}
}

func (f *Forth) Lexer(cmdstr string) Words {
	words := f.NewParser()
	cmd := strings.Split(cmdstr, " ")
	for _, v := range cmd {
		words.parsedwords = append(words.parsedwords, strings.TrimSpace(v))
	}
	return words
}

func (f *Words) GetNextParsedWord() *string {
	if f.parsedIndex >= len(f.parsedwords) {
		return nil
	}
	f.parsedIndex++
	return &f.parsedwords[f.parsedIndex-1]
}

func (words *Words) ExecWord(w string) {
	if words.forth.dictionary[w] == nil {
		i, err := strconv.Atoi(w)
		if err != nil {
			panic(err)
		}
		words.forth.stack.Push(i)
		return
	}
	words.forth.dictionary[w](words)
}

func (w Words) Exec() {
	for {
		v := w.GetNextParsedWord()
		if v == nil {
			return
		}
		w.ExecWord(*v)
	}
}

func NewForthInterpreter() *Forth {
	var f = new(Forth)
	f.dictionary = make(map[string]func(l *Words))

	f.dictionary["+"] = func(l *Words) { f.stack.Push(f.stack.Pop() + f.stack.Pop()) }
	f.dictionary["*"] = func(l *Words) { f.stack.Push(f.stack.Pop() * f.stack.Pop()) }
	f.dictionary["-"] = func(l *Words) { temp := f.stack.Pop(); f.stack.Push(f.stack.Pop() - temp) }
	f.dictionary["/"] = func(l *Words) { temp := f.stack.Pop(); f.stack.Push(f.stack.Pop() / temp) }
	f.dictionary["."] = func(l *Words) { fmt.Println(f.stack.Pop()) }
	f.dictionary["cr"] = func(l *Words) { fmt.Println() }
	f.dictionary["drop"] = func(l *Words) { f.stack.Pop() }
	f.dictionary["dup"] = func(l *Words) { temp := f.stack.Pop(); f.stack.Push(temp); f.stack.Push(temp) }
	f.dictionary["@"] = func(l *Words) { f.stack.Push(f.heap[f.stack.Pop()]) }
	f.dictionary["!"] = func(l *Words) { address := f.stack.Pop(); f.heap[address] = f.stack.Pop() }

	f.dictionary["words"] = func(l *Words) {
		for k := range f.dictionary {
			fmt.Print(k, " ")
		}
		fmt.Println("")
	}

	f.dictionary[".s"] = func(l *Words) {
		fmt.Printf("<%d> ", len(f.stack))
		for _, e := range f.stack {
			fmt.Print(e, " ")
		}
		fmt.Println("")
	}

	f.dictionary["variable"] = func(l *Words) {
		address := len(f.heap)
		f.heap = append(f.heap, 0)
		f.dictionary[*l.GetNextParsedWord()] = func(l *Words) { f.stack.Push(address) }
	}

	f.dictionary[":"] = func(l *Words) {
		funcname := *l.GetNextParsedWord()
		fwords := ExtractFuncFromParser(l)
		f.dictionary[funcname] = func(f *Words) { fwords.Exec() }
	}

	return f
}

func main() {
	f := NewForthInterpreter()

	f.Lexer("variable var1").Exec()
	f.Lexer(": squared dup * ;").Exec()
	f.Lexer("3 squared .").Exec()

	//	f.Parse("2 var1 @ - .")
	//	f.ExecAll()

	//words := f.Lexer("2 3 + .")
	//words.Exec()
	//f.Exec("words")
	//f.Exec("1 2 3 .s")
}
