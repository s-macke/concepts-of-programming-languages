package forth

func (parsedWords parsedWords) compile(forth *Forth) *programState {
	s := &programState{funcs: make([]func(), len(parsedWords))}

	// as default, all words do nothing
	for i := 0; i < len(parsedWords); i++ {
		s.funcs[i] = func() {}
	}

	for i := 0; i < len(parsedWords); i++ {
		switch parsedWords[i].name {

		case "variable":
			address := len(forth.heap)
			forth.heap = append(forth.heap, 0)
			forth.dict[parsedWords[i].parameterString] = func() { forth.stack.Push(address) }

		case "constant":
			parameter := parsedWords[i].parameterString
			s.funcs[i] = func() {
				value := forth.stack.Pop()
				forth.dict[parameter] = func() { forth.stack.Push(value) }
			}

		case ":":
			newFunc := parsedWords[i+1 : parsedWords[i].parameterInt+1].compile(forth)
			forth.dict[parsedWords[i].parameterString] = newFunc.Execute
			i = parsedWords[i].parameterInt // skip the whole function
		case ";": // do nothing

		case "recurse":
			s.funcs[i] = func() { s.instructionPointer = -1 }

		case "if":
			parameter := parsedWords[i].parameterInt
			s.funcs[i] = func() {
				if forth.stack.Pop() == 0 {
					s.instructionPointer += parameter
				}
			}

		case "loop":
			parameter := parsedWords[i].parameterInt
			s.funcs[i] = func() {
				index := forth.returnStack.Pop() + 1
				end := forth.returnStack.Get(0)

				if index < end {
					forth.returnStack.Push(index)
					s.instructionPointer += parameter
				} else {
					forth.returnStack.Pop()
				}
			}

		case "then": // do nothing

		default:
			s.funcs[i] = forth.getWordFromDictionary(parsedWords[i].name)
		}
	}
	return s
}
