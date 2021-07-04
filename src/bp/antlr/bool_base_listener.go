// Code generated from Bool.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Bool

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseBoolListener is a complete listener for a parse tree produced by BoolParser.
type BaseBoolListener struct{}

var _ BoolListener = &BaseBoolListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBoolListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBoolListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBoolListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBoolListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNot is called when production Not is entered.
func (s *BaseBoolListener) EnterNot(ctx *NotContext) {}

// ExitNot is called when production Not is exited.
func (s *BaseBoolListener) ExitNot(ctx *NotContext) {}

// EnterParenthesis is called when production Parenthesis is entered.
func (s *BaseBoolListener) EnterParenthesis(ctx *ParenthesisContext) {}

// ExitParenthesis is called when production Parenthesis is exited.
func (s *BaseBoolListener) ExitParenthesis(ctx *ParenthesisContext) {}

// EnterVariable is called when production Variable is entered.
func (s *BaseBoolListener) EnterVariable(ctx *VariableContext) {}

// ExitVariable is called when production Variable is exited.
func (s *BaseBoolListener) ExitVariable(ctx *VariableContext) {}

// EnterOr is called when production Or is entered.
func (s *BaseBoolListener) EnterOr(ctx *OrContext) {}

// ExitOr is called when production Or is exited.
func (s *BaseBoolListener) ExitOr(ctx *OrContext) {}

// EnterAnd is called when production And is entered.
func (s *BaseBoolListener) EnterAnd(ctx *AndContext) {}

// ExitAnd is called when production And is exited.
func (s *BaseBoolListener) ExitAnd(ctx *AndContext) {}
