package forth

type runContext struct {
	forth              *Forth
	instructionPointer int
	words              []Word
	output             string
}

func (c *runContext) getParameter() int {
	return c.words[c.instructionPointer].parameterInt
}

func (c *runContext) StepOverNextWord() string {
	c.instructionPointer++
	return c.words[c.instructionPointer].name
}

func (c *runContext) Pop() int {
	return c.forth.stack.Pop()
}

func (c *runContext) Push(v int) {
	c.forth.stack.Push(v)
}

func (c *runContext) booleanPush(result bool) {
	if result {
		c.forth.stack.Push(-1)
	} else {
		c.forth.stack.Push(0)
	}
}

func (c *runContext) Run() {
	for {
		if c.instructionPointer >= len(c.words) {
			break
		}
		c.forth.dict.getWord(c.words[c.instructionPointer].name)(c)
		c.instructionPointer++
	}
}

func (c *runContext) Clone() runContext {
	return runContext{
		forth:              c.forth,
		instructionPointer: 0,
		words:              c.words,
		output:             "",
	}
}
