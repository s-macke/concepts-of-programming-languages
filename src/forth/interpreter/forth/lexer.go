package forth

import (
	"fmt"
	"strings"
)

type lexedWords []string

func Lexer(cmd string) lexedWords {
	words := strings.Fields(cmd)
	if debug {
		fmt.Printf("lexer ['%s']\n", strings.Join(words, "', '"))
	}
	return words
}
