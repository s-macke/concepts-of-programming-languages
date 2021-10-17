// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

package parser

import (
	"fmt"
	"testing"

	"github.com/s-macke/concepts-of-programming-languages/src/oop/boolparser/lexer"
)

func TestParser_Eval(t *testing.T) {
	// create parser
	p := NewParser(lexer.NewLexer("a&b&!c"))

	// set vars
	vars := map[string]bool{
		"a": true,
		"b": true,
		"c": false,
	}
	if p.Eval(vars) != true {
		t.Error(fmt.Sprintf("Wrong result detected"))
	}

	// set vars
	vars = map[string]bool{
		"a": true,
		"b": true,
		"c": true,
	}
	if p.Eval(vars) != false {
		t.Error(fmt.Sprintf("Wrong result detected"))
	}

	// set vars
	vars = map[string]bool{
		"a": false,
		"b": false,
		"c": false,
	}
	if p.Eval(vars) != false {
		t.Error(fmt.Sprintf("Wrong result detected"))
	}

	p = NewParser(lexer.NewLexer("a & (b | c & b) & d"))

	// set vars
	vars = map[string]bool{
		"a": true,
		"b": true,
		"c": false,
		"d": true,
	}
	if p.Eval(vars) != true {
		t.Error(fmt.Sprintf("Wrong result detected"))
	}

	// test string support
	if p.String() != "&(&('a',|('b',&('c','b'))),'d')" {
		t.Error(fmt.Sprintf("Wrong string representation: %v", p.String()))
	}
}
