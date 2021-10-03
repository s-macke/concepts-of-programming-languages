package main

import (
	"fmt"
	"strings"
)

type LexedWords []string

func Lexer(cmd string) LexedWords {
	words := strings.Fields(cmd)
	if debug {
		fmt.Printf("Lexer %v\n", words)
	}
	return words
}
