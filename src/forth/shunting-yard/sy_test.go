package main

import (
	"fmt"
	"testing"
)

type TestList struct {
	input    string
	expected string
	result   int
}

var tests = []TestList{
	{"", "[]", 0},
	{"3 + 4", "[3 4 +]", 7},
	{"1 - 2", "[1 2 -]", -1},
	{"6 / 2", "[6 2 /]", 3},
	{"3 + 4 * 5", "[3 4 5 * +]", 23},
	{"3 + 4 * 2 * ( 5 + 6 )", "[3 4 2 * 5 6 + * +]", 91},
}

func TestAll(t *testing.T) {
	for _, test := range tests {
		fmt.Println(test)
		postfix := Lexer(test.input).ToPostfix()
		if fmt.Sprintf("%v", postfix) != test.expected {
			t.Errorf("Expected %s, got %v", test.expected, postfix)
		}
		if test.expected == "[]" {
			continue
		}
		result, err := postfix.Evaluate()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if result != test.result {
			t.Errorf("Expected %d, got %d", test.result, result)
		}
	}
}
