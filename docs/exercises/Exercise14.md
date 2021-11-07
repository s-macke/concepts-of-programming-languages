# Exercise 14.1 - Forth

----
Experiment with Forth

Write the following functions:

----

## Exercise 14.2 - Lexer, Parser and Interpreter for Forth

**Disclaimer:** Feel free you use your very own software design.

### Runtime 
Write a Forth runtime by implementing
- a stack for integers
- a dictionary which maps strings to functions
- a heap memory to store variables

Implement the following Forth words 
"+", "-", ".", "dup" 

### Lexer
Write a lexer for the Forth programming language.
Use the strings.Fields function


### Simple Interpreter

Write an interpreter that 
- searches each word in the dictionary and executes it.
- otherwise tries to intepret the string as number
- fails if none of the above applies

🤥 **Write tests!** 🤥


### Parser

In Forth the parser combines control structure words 



The parser is split up into two parts:

- Define and implement an Abstract Syntax Tree
- Token parsing and Abstract Syntax Tree building

_Reminder:_ Use the following grammar definition:

```bnf
<expression> ::= <term> { <or> <term> }
<term> ::= <factor> { <and> <factor> }
<factor> ::= <var> | <not> <factor> | (<expression>)
<or>  ::= '|'
<and> ::= '&'
<not> ::= '!'
<var> ::= '[a-zA-Z0-9]+'
```

### Exercise 4.2.1 - Abstract Syntax Tree (AST)

Write a program which builds an AST with nodes to evaluate logical expressions with (And, Not, Or with variables).

```text
Sample Expression: `A AND B OR C`

             ----------
             |   OR   |
             ----------
            /          \
        ---------      ----------
        |  AND  |      |  Var:C |
        ---------      ----------
        /       \
  ---------   ---------
  | Var:A |   | Var:B |
  ---------   ---------
```

The tree should be evaluated with a evaluation methods which supports named variables:

```go
eval(vars map[string]bool) bool
```

_Why named variables:_ This allows us to build the AST once and use it for multiple variable values.

Notes that might help:

- Interfaces and Polymorphism
- Nodes are different but what are the commonalities?
- Simply follow the rules

🤥 **Write tests! Otherwise it does not happen!** 🤥

Write a unit test which builds the AST and evaluate the expression with given boolean values for the variables A, B, C.

### Exercise 4.2.2 Recursive Descent Parser

Write a recursive descent parser. The parser must implement the grammar rules  (that was enough hint).

🤥 **Write tests! Otherwise it does not happen!** 🤥

### Exercise 4.3 Antlr

We now use Antlr to generate a lexer and a parser for a given grammar definition.

Follow the go Antlr quick-start: <https://github.com/antlr/antlr4/blob/master/doc/go-target.md>

You need to do the following things:

- Antlr Setup (see above)
- Define an Antlr grammar file (`boolparser.g4`)
- Generate lexer and parser source code
- Use the generated files to parse boolean expressions

Should be not to hard 🤙
