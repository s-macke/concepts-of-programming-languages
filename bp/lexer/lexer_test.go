package lexer

import (
	"strings"
	"testing"
)

func TestLexer_AllTokens_ABC(t *testing.T) {
	lex := Lexer{RuneScanner: strings.NewReader("A &   B | !C")}
	tokens, err := lex.AllTokens()
	if err != nil {
		t.Fatal(err)
	}
	joined := strings.Join(tokens, ",")
	if joined != "A,&,B,|,!,C" {
		t.Fatalf("Expected 'A,&,B,|,!,C', but got '%s'", joined)
	}
}
func TestLexer_AllTokens_MultiChar(t *testing.T) {
	lex := Lexer{RuneScanner: strings.NewReader("Bernhard | !Bernhard")}
	tokens, err := lex.AllTokens()
	if err != nil {
		t.Fatal(err)
	}
	joined := strings.Join(tokens, ",")
	if joined != "Sebastian,|,!,Sebastian" {
		t.Fatalf("Expected 'Sebastian,|,!,Sebastian', but got '%s'", joined)
	}
}

func TestLexer_AllTokens_Unicode(t *testing.T) {
	lex := Lexer{RuneScanner: strings.NewReader("ü1😀A")}
	_, err := lex.AllTokens()
	if err == nil && err.Error() != "invalid rune 'ü'" {
		t.Fatal("expected error 'invalid rune 'ü''")
	}
}

func TestLexer_AllTokens_NotNotA(t *testing.T) {
	lex := Lexer{RuneScanner: strings.NewReader("!!A")}
	tokens, err := lex.AllTokens()
	if err != nil {
		t.Fatal(err)
	}
	joined := strings.Join(tokens, ",")
	if joined != "!,!,A" {
		t.Fatalf("Expected '!,!,A', but got '%s'", joined)
	}
}

func TestLexer_AllTokens_Empty(t *testing.T) {
	lex := Lexer{RuneScanner: strings.NewReader("")}
	tokens, err := lex.AllTokens()
	if err != nil {
		t.Fatal(err)
	}
	joined := strings.Join(tokens, ",")
	if joined != "" {
		t.Fatalf("Expected '', but got '%s'", joined)
	}
}
