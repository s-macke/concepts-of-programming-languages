package main

import (
	"fmt"
	antp "github.com/0xqab/concepts-of-programming-languages/bp/antlr"
	"github.com/0xqab/concepts-of-programming-languages/bp/lexer"
	"github.com/0xqab/concepts-of-programming-languages/bp/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strings"
)

func main() {
	source := "A & B | !C"
	reader := strings.NewReader(source)
	lex := lexer.Lexer{RuneScanner: reader}
	pars := parser.Parser{Tokenizer: lex}
	ast, err := pars.Parse()
	if err != nil {
		panic(err)
	}
	fmt.Printf("AST: %v\n", ast)
	fmt.Printf("eval(true, true, true) = %t\n", ast.Eval(map[string]bool{"A": true, "B": true, "C": true}))
	fmt.Printf("eval(true, true, false) = %t\n", ast.Eval(map[string]bool{"A": true, "B": true, "C": false}))
	fmt.Printf("eval(true, false, true) = %t\n", ast.Eval(map[string]bool{"A": true, "B": false, "C": true}))
	fmt.Printf("eval(false, true, true) = %t\n", ast.Eval(map[string]bool{"A": false, "B": true, "C": true}))
	fmt.Printf("eval(false, false, false) = %t\n", ast.Eval(map[string]bool{"A": false, "B": false, "C": false}))
	fmt.Printf("eval(false, false, true) = %t\n", ast.Eval(map[string]bool{"A": false, "B": false, "C": true}))

	antlrLexer := antp.NewBoolLexer(antlr.NewInputStream(source))
	antlrParser := antp.NewBoolParser(antlr.NewCommonTokenStream(antlrLexer, 0))
	antlrAst := antlrParser.Expr()
	fmt.Printf("antlr.eval(true, true, true) = %t\n", antlrEval(antlrAst, map[string]bool{"A": true, "B": true, "C": true}))
	fmt.Printf("antlr.eval(true, true, false) = %t\n", antlrEval(antlrAst, map[string]bool{"A": true, "B": true, "C": false}))
	fmt.Printf("antlr.eval(true, false, true) = %t\n", antlrEval(antlrAst, map[string]bool{"A": true, "B": false, "C": true}))
	fmt.Printf("antlr.eval(false, true, true) = %t\n", antlrEval(antlrAst, map[string]bool{"A": false, "B": true, "C": true}))
	fmt.Printf("antlr.eval(false, false, false) = %t\n", antlrEval(antlrAst, map[string]bool{"A": false, "B": false, "C": false}))
	fmt.Printf("antlr.eval(false, false, true) = %t\n", antlrEval(antlrAst, map[string]bool{"A": false, "B": false, "C": true}))
}

func antlrEval(antlrAst antp.IExprContext, vars map[string]bool) bool {
	evaluator := antp.NewEvaluator(vars)
	antlr.ParseTreeWalkerDefault.Walk(evaluator, antlrAst)
	return evaluator.Result()
}
