package parser

type Evaluator struct {
	*BaseBoolListener
	vars  map[string]bool
	stack []bool
}

func NewEvaluator(vars map[string]bool) *Evaluator {
	return &Evaluator{
		BaseBoolListener: &BaseBoolListener{},
		vars:             vars,
		stack:            []bool{},
	}
}

func (e *Evaluator) push(b bool) {
	e.stack = append(e.stack, b)
}

func (e *Evaluator) pop() bool {
	if len(e.stack) < 1 {
		panic("empty stack")
	}
	result := e.stack[len(e.stack)-1]
	e.stack = e.stack[:len(e.stack)-1]
	return result
}

func (e *Evaluator) Result() bool {
	return e.pop()
}

func (e *Evaluator) ExitNot(ctx *NotContext) {
	e.push(!e.pop())
}

func (e *Evaluator) ExitAnd(ctx *AndContext) {
	e.push(e.pop() && e.pop())
}

func (e *Evaluator) ExitOr(ctx *OrContext) {
	e.push(e.pop() || e.pop())
}

func (e *Evaluator) ExitVariable(ctx *VariableContext) {
	e.push(e.vars[ctx.GetText()])
}
