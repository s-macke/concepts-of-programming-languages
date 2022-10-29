package main

import (
	"fmt"
	"testing"
)

type TestList struct {
	input    string
	expected string
}

var tests = []TestList{
	{"", "[]"},
	{"3 + 4", "[3 4 +]"},
	{"3 + 4 * 5", "[3 4 5 * +]"},
	{"3 + 4 * 2 * ( 5 + 6 )", "[3 4 2 * 5 6 + * +]"},
}

func TestAll(t *testing.T) {
	for _, test := range tests {
		sy := NewShuntingYard(test.input)
		sy.Lex()
		sy.Parse()
		if fmt.Sprintf("%v", sy.output) != test.expected {
			t.Errorf("Expected %s, got %v", test.expected, sy.output)
		}
	}
}
