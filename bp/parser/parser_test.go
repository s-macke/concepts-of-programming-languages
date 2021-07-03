package parser

import (
	"github.com/s-macke/concepts-of-programming-languages/bp/lexer"
	"strings"
	"testing"
)

func TestSimpleExample(t *testing.T) {
	source := "A & B | !C"
	reader := strings.NewReader(source)
	lex := lexer.Lexer{RuneScanner: reader}
	pars := Parser{Tokenizer: lex}
	ast, err := pars.Parse()
	if err != nil {
		panic(err)
	}

	if ast.Eval(map[string]bool{"A": true, "B": true, "C": true}) != true {
		t.Error("eval(true, true, true) should be true, but was false")
	}
	if ast.Eval(map[string]bool{"A": true, "B": true, "C": false}) != true {
		t.Error("eval(true, true, false) should be true, but was false")
	}
	if ast.Eval(map[string]bool{"A": true, "B": false, "C": true}) != false {
		t.Error("eval(true, false, true) should be false, but was true")
	}
	if ast.Eval(map[string]bool{"A": false, "B": true, "C": true}) != false {
		t.Error("eval(false, true, true) should be false, but was true")
	}
	if ast.Eval(map[string]bool{"A": false, "B": false, "C": false}) != true {
		t.Error("eval(false, false, false) should be true, but was false")
	}
	if ast.Eval(map[string]bool{"A": false, "B": false, "C": true}) != false {
		t.Error("eval(false, false, true) should be false, but was true")
	}
}

func TestParenthesisExample(t *testing.T) {
	source := "A & (B | C)"
	reader := strings.NewReader(source)
	lex := lexer.Lexer{RuneScanner: reader}
	pars := Parser{Tokenizer: lex}
	ast, err := pars.Parse()
	if err != nil {
		panic(err)
	}

	if ast.Eval(map[string]bool{"A": true, "B": true, "C": true}) != true {
		t.Error("eval(true, true, true) should be true, but was false")
	}
	if ast.Eval(map[string]bool{"A": true, "B": true, "C": false}) != true {
		t.Error("eval(true, true, false) should be true, but was false")
	}
	if ast.Eval(map[string]bool{"A": true, "B": false, "C": true}) != true {
		t.Error("eval(true, false, true) should be true, but was false")
	}
	if ast.Eval(map[string]bool{"A": false, "B": true, "C": true}) != false {
		t.Error("eval(false, true, true) should be false, but was true")
	}
	if ast.Eval(map[string]bool{"A": false, "B": false, "C": false}) != false {
		t.Error("eval(false, false, false) should be false, but was true")
	}
}

func TestNotNotExample(t *testing.T) {
	source := "!!A"
	reader := strings.NewReader(source)
	lex := lexer.Lexer{RuneScanner: reader}
	pars := Parser{Tokenizer: lex}
	ast, err := pars.Parse()
	if err != nil {
		panic(err)
	}

	if ast.Eval(map[string]bool{"A": true}) != true {
		t.Error("eval(true) should be true, but was false")
	}
	if ast.Eval(map[string]bool{"A": false}) != false {
		t.Error("eval(true) should be false, but was true")
	}
}
