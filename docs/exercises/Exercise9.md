# Exercise 9.1 - Forth

- Write a function, which takes one integer as argument and calculates the cube x³ of that number
- Write a function which which prints the numbers from 1 to 10
- Write a function, which checks whether the given number on the stack is a prime number (loop over all odd numbers and test via the word "mod")
- Write a function "swapv", which swaps the content of two variables

```Forth
  variable a   1 a !   ( a = 1 )
  variable b   2 b !   ( b = 2 )
  a b swapv
  a @ .  ( outputs 2 )
  b @ .  ( outputs 1 )
```

## Exercise 9.2 - Interpreter for Forth

**Disclaimer:** Feel free you use your very own software design.

Ignore Forth strings for this exercise.

Write a Forth runtime by implementing
- a stack for integers
- a dictionary which maps strings to functions

🤥 **Write tests!** 🤥

Write an interpreter that
- searches each word in the dictionary and executes it.
- otherwise tries to intepret the string as number
- fails if none of the above applies

Implement at least the following Forth words 
  "+", "-", ".", "dup", "drop", "words" 

### Constants

- Allow a Forth word to jump over the next word and to receive the word string
- Implement the constant word

### Variables

- Implement a heap memory to store variables
- Implement the variable word
- Add the forth words @ and !

### If/then

- Implement the words "true" and "false"
- Allow a lexed word to have an additional jumpTo parameter
- Implement a parser to determine the word pairs as discussed in the lecture

### Implement comments

- Implement the comment words "(" and ")"

### Have fun

Continue just as you like

- Implement do/loop. You have to add an additional "return stack" to store the meta information for the loop.
- Implement functions. Are you able to execute them recursively?

Link to small code snipplets:
https://rosettacode.org/wiki/Category:Forth

