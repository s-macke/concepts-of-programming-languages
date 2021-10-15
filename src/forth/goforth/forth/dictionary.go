package forth

import (
	"fmt"
	"strconv"
)

type dictionary map[string]func() // dictionary

func (f *Forth) fillDictionary() {
	f.dict["+"] = func() { f.stack.Push(f.stack.Pop() + f.stack.Pop()) }
	f.dict["*"] = func() { f.stack.Push(f.stack.Pop() * f.stack.Pop()) }
	f.dict["-"] = func() { temp := f.stack.Pop(); f.stack.Push(f.stack.Pop() - temp) }
	f.dict["/"] = func() { temp := f.stack.Pop(); f.stack.Push(f.stack.Pop() / temp) }

	f.dict[">"] = func() { temp := f.stack.Pop(); f.booleanPush(f.stack.Pop() > temp) }
	f.dict["<"] = func() { temp := f.stack.Pop(); f.booleanPush(f.stack.Pop() < temp) }
	f.dict["="] = func() { f.booleanPush(f.stack.Pop() == f.stack.Pop()) }
	f.dict["<>"] = func() { f.booleanPush(f.stack.Pop() == f.stack.Pop()) }

	f.dict["."] = func() { f.output += fmt.Sprintf("%d ", f.stack.Pop()) }
	f.dict["cr"] = func() { f.output += fmt.Sprintln() }
	f.dict["drop"] = func() { f.stack.Pop() }
	f.dict["dup"] = func() { temp := f.stack.Pop(); f.stack.Push(temp); f.stack.Push(temp) }
	f.dict["swap"] = func() { temp, temp2 := f.stack.Pop(), f.stack.Pop(); f.stack.Push(temp); f.stack.Push(temp2) }
	f.dict["@"] = func() { f.stack.Push(f.heap[f.stack.Pop()]) }
	f.dict["!"] = func() { address := f.stack.Pop(); f.heap[address] = f.stack.Pop() }

	f.dict[">r"] = func() { f.returnStack.Push(f.stack.Pop()) }
	f.dict["r>"] = func() { f.stack.Push(f.returnStack.Pop()) }
	f.dict["i"] = func() { f.stack.Push(f.returnStack.Get(0)) }
	f.dict["do"] = func() { f.dict["swap"](); f.dict[">r"](); f.dict[">r"]() }

	f.dict["bye"] = func() { panic("Program stopped") }

	f.dict["words"] = func() {
		for k := range f.dict {
			f.output += fmt.Sprint(k, " ")
		}
		f.output += fmt.Sprintln("")
	}

	f.dict[".s"] = func() {
		f.output += fmt.Sprintf("<%d> ", len(f.stack))
		for _, e := range f.stack {
			f.output += fmt.Sprint(e, " ")
		}
	}
}

func (f *Forth) getWordFromDictionary(w string) func() {
	if f.dict[w] == nil {
		i, err := strconv.Atoi(w)
		if err != nil {
			panic("Word '" + w + "' is neither a valid word nor a number")
		}
		return func() {
			f.stack.Push(i)
		}
	}
	return f.dict[w]
}
