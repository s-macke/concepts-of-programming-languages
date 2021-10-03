package main

import "fmt"

type ParsedWord struct {
	name            string
	parameterInt    int
	parameterString string
}

type ParsedWords []ParsedWord

func (lexedWords LexedWords) Parser() ParsedWords {
	var words ParsedWords

	var controlStructureStack Stack

	parsedId := 0
	for i := 0; i < len(lexedWords); i++ {
		word := ParsedWord{name: lexedWords[i], parameterInt: -1, parameterString: ""}

		switch lexedWords[i] {

		case "(":
			for ; lexedWords[i] != ")"; i++ {
				fmt.Println(lexedWords[i])
			}
			continue

		case "do":
			controlStructureStack.Push(parsedId)

		case "loop":
			id := controlStructureStack.Pop()
			if words[id].name != "do" {
				panic("loop without do")
			}
			word.parameterInt = id

		case "if":
			controlStructureStack.Push(parsedId)

		case "then":
			id := controlStructureStack.Pop()
			if words[id].name != "if" {
				panic("then without if")
			}
			words[id].parameterInt = parsedId

		case "variable", "constant":
			word.parameterString = lexedWords[i+1]
			i++

		case ":":
			word.parameterString = lexedWords[i+1]
			i++
			controlStructureStack.Push(parsedId)

		case ";":
			id := controlStructureStack.Pop()
			if words[id].name != ":" {
				panic("function start not found")
			}
			words[id].parameterInt = parsedId - 1
		}

		words = append(words, word)
		parsedId++
	}

	if debug {
		for i := 0; i < len(words); i++ {
			fmt.Printf("Parser %2d: word=%10s   param=%3d   param=%s\n", i, words[i].name, words[i].parameterInt, words[i].parameterString)
		}
	}
	return words
}
