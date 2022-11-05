package main

import (
	"fmt"
	"strconv"
	"strings"
)

var precedence = map[string]int{
	"*": 2,
	"/": 2,
	"+": 1,
	"-": 1,
}

type Tokens []string

func Lexer(input string) Tokens {
	return strings.Fields(input)
}

type Postfix []string

// Parse the input expression by performing the shunting yard algorithm
func (tokens Tokens) ToPostfix() Postfix {
	operatorStack := NewStack[string]()
	var output []string

	for _, token := range tokens {
		//fmt.Println("Parse ", token)

		switch token {
		case "+", "-", "*", "/":
			for !operatorStack.IsEmpty() { // Check precedence of operators
				op := operatorStack.Get(0)
				if precedence[op] < precedence[token] {
					break
				}
				output = append(output, operatorStack.Pop())
			}
			operatorStack.Push(token)
		case "(":
			operatorStack.Push(token)

		case ")": // Highest precedence
			for !operatorStack.IsEmpty() {
				op := operatorStack.Pop()
				if op == "(" {
					break
				}
				output = append(output, op)
			}

		default: // A number or variable
			output = append(output, token)
		}
		//fmt.Println("output:", output, "Operator Stack", operatorStack.data)
	}

	// Pop all operators from the stack and append them to the output queue.
	for !operatorStack.IsEmpty() {
		output = append(output, operatorStack.Pop())
	}
	return output
}

func (pf Postfix) Evaluate() (int, error) {
	s := NewStack[int]()
	for _, expr := range pf {
		switch expr {
		case "+":
			s.Push(s.Pop() + s.Pop())
		case "-":
			right := s.Pop()
			s.Push(s.Pop() - right)
		case "*":
			s.Push(s.Pop() * s.Pop())
		case "/":
			right := s.Pop()
			s.Push(s.Pop() / right)
		default:
			v, err := strconv.Atoi(expr)
			if err != nil {
				return 0, err
			}
			s.Push(v)
		}
	}
	return s.Pop(), nil
}

func main() {
	result, err := Lexer("3 + 4 * 2 * ( 5 + 6 )").ToPostfix().Evaluate()
	if err != nil {
		panic(err)
	}
	fmt.Println("Result:", result)
}
