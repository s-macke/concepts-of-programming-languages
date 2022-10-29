package main

import (
	"fmt"
	"strings"
)

// Implement the shunting-yard algorithm for parsing infix expressions.
type ShuntingYard struct {
	// The input expression
	input string
	// Input Tokens
	tokens []string
	// The output queue
	output []string
	// The operator stack
	operatorStack *Stack[string]

	precedence map[string]int // The precedence of operators

}

// Create a new ShuntingYard instance.
func NewShuntingYard(input string) *ShuntingYard {
	return &ShuntingYard{
		input:         input,
		output:        make([]string, 0),
		operatorStack: NewStack[string](),
		precedence: map[string]int{
			"(": 1,
			"*": 3,
			"/": 3,
			"+": 2,
			"-": 2,
		},
	}
}

func (sy *ShuntingYard) Lex() []string {
	sy.tokens = strings.Fields(sy.input)
	return sy.tokens
}

// Parse the input expression.
func (sy *ShuntingYard) Parse() []string {
	for _, token := range sy.tokens {
		fmt.Println("Parse ", token)

		switch token {
		case "+", "-", "*", "/":
			for !sy.operatorStack.IsEmpty() { // Check precedence of operators
				op := sy.operatorStack.Get(0)
				if sy.precedence[op] < sy.precedence[token] {
					break
				}
				sy.output = append(sy.output, sy.operatorStack.Pop())
			}
			sy.operatorStack.Push(token)
		case "(":
			sy.operatorStack.Push(token)

		case ")": // Highest precedence
			for !sy.operatorStack.IsEmpty() {
				op := sy.operatorStack.Pop()
				if op == "(" {
					break
				}
				sy.output = append(sy.output, op)
			}

		default: // A number or variable
			sy.output = append(sy.output, token)
		}
		fmt.Println("output:", sy.output, "Operator Stack", sy.operatorStack.data)
	}

	// Pop all operators from the stack and append them to the output queue.
	for !sy.operatorStack.IsEmpty() {
		sy.output = append(sy.output, sy.operatorStack.Pop())
	}

	return sy.output
}

func main() {
	//sy := NewShuntingYard("3 + 4 * 2 / ( 1 - 5 ) ^ 2 ^ 3")
	sy := NewShuntingYard("3 + 4 * 2 * ( 5 + 6 )")
	sy.Lex()
	fmt.Println("Lexer:", sy.tokens)
	sy.Parse()
	fmt.Println("Result:", sy.output)
}
