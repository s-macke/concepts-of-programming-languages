package forth

import "fmt"

type parsedWord struct {
	name            string
	parameterInt    int
	parameterString string
}

type parsedWords []parsedWord

func (lexedWords lexedWords) Parser() parsedWords {
	var words parsedWords

	var controlStructureStack stack

	parsedID := 0
	for i := 0; i < len(lexedWords); i++ {
		word := parsedWord{name: lexedWords[i], parameterInt: 0, parameterString: ""}

		switch lexedWords[i] {

		case "(":
			for ; lexedWords[i] != ")"; i++ {
			}
			continue

		case "do":
			controlStructureStack.Push(parsedID)

		case "loop":
			id := controlStructureStack.Pop()
			if words[id].name != "do" {
				panic("loop without do")
			}
			word.parameterInt = id - parsedID

		case "if":
			controlStructureStack.Push(parsedID)

		case "then":
			id := controlStructureStack.Pop()
			if words[id].name != "if" {
				panic("then without if")
			}
			words[id].parameterInt = parsedID - id

		case "variable", "constant":
			word.parameterString = lexedWords[i+1]
			i++

		case ":":
			word.parameterString = lexedWords[i+1]
			i++
			controlStructureStack.Push(parsedID)

		case ";":
			id := controlStructureStack.Pop()
			if words[id].name != ":" {
				panic("function start not found")
			}
			words[id].parameterInt = parsedID - 1
		}

		words = append(words, word)
		parsedID++
	}

	if debug {
		for i := 0; i < len(words); i++ {
			fmt.Printf("Parser %2d: word=%10s   param=%3d   param=%s\n", i, words[i].name, words[i].parameterInt, words[i].parameterString)
		}
	}
	return words
}
