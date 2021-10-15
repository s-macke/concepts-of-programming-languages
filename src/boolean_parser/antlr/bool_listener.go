// Code generated from Bool.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Bool

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BoolListener is a complete listener for a parse tree produced by BoolParser.
type BoolListener interface {
	antlr.ParseTreeListener

	// EnterNot is called when entering the Not production.
	EnterNot(c *NotContext)

	// EnterParenthesis is called when entering the Parenthesis production.
	EnterParenthesis(c *ParenthesisContext)

	// EnterVariable is called when entering the Variable production.
	EnterVariable(c *VariableContext)

	// EnterOr is called when entering the Or production.
	EnterOr(c *OrContext)

	// EnterAnd is called when entering the And production.
	EnterAnd(c *AndContext)

	// ExitNot is called when exiting the Not production.
	ExitNot(c *NotContext)

	// ExitParenthesis is called when exiting the Parenthesis production.
	ExitParenthesis(c *ParenthesisContext)

	// ExitVariable is called when exiting the Variable production.
	ExitVariable(c *VariableContext)

	// ExitOr is called when exiting the Or production.
	ExitOr(c *OrContext)

	// ExitAnd is called when exiting the And production.
	ExitAnd(c *AndContext)
}
