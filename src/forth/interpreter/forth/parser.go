package forth

import "fmt"

type Word struct {
	name         string
	parameterInt int
}

type Parser []Word

func NewParser(lexedWords lexedWords) Parser {
	var words Parser
	for i := 0; i < len(lexedWords); i++ {
		word := Word{name: lexedWords[i], parameterInt: 0}
		words = append(words, word)
	}
	return words
}

func (parser Parser) Parse() {
	var controlStructureStack stack

	for i := 0; i < len(parser); i++ {
		switch parser[i].name {
		case "(":
			commentStartPosition := i
			for ; parser[i].name != ")"; i++ {
			}
			parser[commentStartPosition].parameterInt = i
			continue

		case "do", "if", ":":
			controlStructureStack.Push(i)

		case "loop":
			id := controlStructureStack.Pop()
			if parser[id].name != "do" {
				panic("loop without do")
			}
			parser[i].parameterInt = id

		case "then", "else":
			id := controlStructureStack.Pop()
			if parser[id].name != "if" && parser[id].name != "else" {
				panic("then without if or else")
			}
			parser[id].parameterInt = i
			if parser[i].name == "else" {
				controlStructureStack.Push(i)
			}

		case ";":
			id := controlStructureStack.Pop()
			if parser[id].name != ":" {
				panic("function start not found")
			}
			parser[id].parameterInt = i
		}
	}

	if debug {
		for i := 0; i < len(parser); i++ {
			fmt.Printf("Parser %2d: word=%10s   param=%3d\n", i, parser[i].name, parser[i].parameterInt)
		}
	}
}
