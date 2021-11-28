package forth

import (
	"fmt"
	"strconv"
)

type dictFunc func(c *runContext)

type dictionary map[string]dictFunc // dictionary

func NewDictionary() dictionary {
	dict := make(dictionary)
	dict.fill()
	return dict
}

func (dict dictionary) fill() {
	dict["+"] = func(c *runContext) { c.Push(c.Pop() + c.Pop()) }
	dict["*"] = func(c *runContext) { c.Push(c.Pop() * c.Pop()) }
	dict["-"] = func(c *runContext) { temp := c.Pop(); c.Push(c.Pop() - temp) }
	dict["/"] = func(c *runContext) { temp := c.Pop(); c.Push(c.Pop() / temp) }

	dict["true"] = func(c *runContext) { c.Push(-1) }
	dict["false"] = func(c *runContext) { c.Push(0) }
	dict[">"] = func(c *runContext) { temp := c.Pop(); c.booleanPush(c.Pop() > temp) }
	dict["<"] = func(c *runContext) { temp := c.Pop(); c.booleanPush(c.Pop() < temp) }
	dict["="] = func(c *runContext) { c.booleanPush(c.Pop() == c.Pop()) }
	dict["<>"] = func(c *runContext) { c.booleanPush(c.Pop() == c.Pop()) }
	dict["0>"] = func(c *runContext) { c.booleanPush(c.Pop() > 0) }
	dict["1-"] = func(c *runContext) { c.Push(c.Pop() - 1) }

	dict["."] = func(c *runContext) { c.output += fmt.Sprintf("%d ", c.Pop()) }
	dict["cr"] = func(c *runContext) { c.output += fmt.Sprintln() }
	dict["drop"] = func(c *runContext) { c.Pop() }
	dict["dup"] = func(c *runContext) { temp := c.Pop(); c.Push(temp); c.Push(temp) }
	dict["swap"] = func(c *runContext) { temp, temp2 := c.Pop(), c.Pop(); c.Push(temp); c.Push(temp2) }
	dict["@"] = func(c *runContext) { c.Push(c.forth.heap[c.Pop()]) }
	dict["!"] = func(c *runContext) { address := c.Pop(); c.forth.heap[address] = c.Pop() }

	dict[">r"] = func(c *runContext) { c.forth.returnStack.Push(c.Pop()) }
	dict["r>"] = func(c *runContext) { c.Push(c.forth.returnStack.Pop()) }
	dict["i"] = func(c *runContext) { c.Push(c.forth.returnStack.Get(0)) }
	dict["do"] = func(c *runContext) { c.forth.dict["swap"](c); c.forth.dict[">r"](c); c.forth.dict[">r"](c) }
	dict["loop"] = func(c *runContext) {
		index := c.forth.returnStack.Pop() + 1
		end := c.forth.returnStack.Get(0)

		if index < end {
			c.forth.returnStack.Push(index)
			c.instructionPointer = c.getParameter()
		} else {
			c.forth.returnStack.Pop()
		}
	}

	dict["("] = func(c *runContext) {
		c.instructionPointer = c.getParameter()
	}

	dict["bye"] = func(c *runContext) { panic("Program stopped") }

	dict["constant"] = func(c *runContext) {
		name := c.StepOverNextWord()
		v := c.Pop()
		dict[name] = func(c *runContext) { c.Push(v) }
	}

	dict["variable"] = func(c *runContext) {
		name := c.StepOverNextWord()
		address := len(c.forth.heap)
		c.forth.heap = append(c.forth.heap, 0)
		dict[name] = func(c *runContext) { c.Push(address) }
	}

	dict["words"] = func(c *runContext) {
		for k := range dict {
			c.output += fmt.Sprint(k, " ")
		}
		c.output += fmt.Sprintln("")
	}

	dict["then"] = func(c *runContext) {}
	dict["if"] = func(c *runContext) {
		if c.Pop() == 0 {
			c.instructionPointer = c.getParameter()
		}
	}
	dict["else"] = func(c *runContext) {
		c.instructionPointer = c.getParameter()
	}

	dict["recurse"] = func(c *runContext) {
		ctx := c.Clone()
		ctx.Run()
		c.output += ctx.output
	}

	dict[":"] = func(c *runContext) {
		funcEnd := c.getParameter()
		funcWords := Parser(c.words[c.instructionPointer+2 : funcEnd])
		funcWords.Parse()
		name := c.StepOverNextWord()
		funcContext := c.forth.NewRunContext(funcWords)
		dict[name] = func(c *runContext) {
			ctx := funcContext.Clone()
			ctx.Run()
			c.output += ctx.output
		}
		c.instructionPointer = funcEnd
	}

	dict[".s"] = func(c *runContext) {
		c.output += fmt.Sprintf("<%d> ", len(c.forth.stack))
		for _, e := range c.forth.stack {
			c.output += fmt.Sprint(e, " ")
		}
	}
}

func (dict dictionary) getWord(w string) dictFunc {
	if debug {
		fmt.Println("Exec Word:", w)
	}
	if dict[w] == nil {
		i, err := strconv.Atoi(w)
		if err != nil {
			panic("Word '" + w + "' is neither a valid word nor a number")
		}
		return func(c *runContext) {
			c.Push(i)
		}
	}
	return dict[w]
}
