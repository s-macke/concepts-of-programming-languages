# Exercise - Forth

## Exercise 1 - Forth

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

## Exercise 2 - Implement Shunting Yard Algorithm

Implement the Shunting Yard Algorithm
- Take a string as input such as "1 + 2 * 3". Allowed operators are `+`, `-`, `*`, `/`
- Implement a tokenizer which splits the string into tokens by using the `strings.Fields` function.
- Implement a parser which converts the tokens into a postfix notation.
- Implement a function which evaluates the postfix notation.
- Implement brackets `(`, `)`
 


