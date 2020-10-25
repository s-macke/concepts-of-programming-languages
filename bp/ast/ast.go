package ast

import "fmt"

type Node interface {
	Eval(vars map[string]bool) bool
}

type Not struct {
	Expr Node
}

func (n Not) Eval(vars map[string]bool) bool {
	return !n.Expr.Eval(vars)
}

func (n Not) String() string {
	return fmt.Sprintf("!(%v)", n.Expr)
}

type And struct {
	LeftExpr  Node
	RightExpr Node
}

func (n And) Eval(vars map[string]bool) bool {
	return n.LeftExpr.Eval(vars) && n.RightExpr.Eval(vars)
}

func (n And) String() string {
	return fmt.Sprintf("(%v & %v)", n.LeftExpr, n.RightExpr)
}

type Or struct {
	LeftExpr  Node
	RightExpr Node
}

func (n Or) Eval(vars map[string]bool) bool {
	return n.LeftExpr.Eval(vars) || n.RightExpr.Eval(vars)
}

func (n Or) String() string {
	return fmt.Sprintf("(%v | %v)", n.LeftExpr, n.RightExpr)
}

type Var struct {
	Name string
}

func (n Var) Eval(vars map[string]bool) bool {
	return vars[n.Name]
}

func (n Var) String() string {
	return fmt.Sprintf("%s", n.Name)
}
