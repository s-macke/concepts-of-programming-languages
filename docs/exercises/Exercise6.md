# Exercise 6 - Haskell

Use an online Haskell interpreter such as.

* https://replit.com/languages/haskell
* https://www.jdoodle.com/execute-haskell-online/
* https://www.onlinegdb.com/online_haskell_compiler

Solve the following exercises.

### Exercise 6.1 - Guards
- A signum function returns 1 if the number is positive, -1 if the number is negative and 0 if the number is 0. 
  Write a function with guards that returns the signum of a number.

### Exercise 6.2 - Recursion with pattern matching

- Based on the Factorial example, write a recursive function which calculates the Fibonacci number for a given number n.
  Fibonacci sequence is a sequence of numbers where the next number is the sum of the previous two numbers. 
  The first two numbers in the sequence are 0 and 1. 
  The first 10 numbers in the sequence are 0, 1, 1, 2, 3, 5, 8, 13, 21, 34.

### Exercise 6.3 - Lists

- Implement a recursive sum function which takes a list of numbers and returns the sum of all numbers in the list.
 
```haskell
  main = print $ sumList [1,2,3]  -- returns 6
```

- Implement a string reverser

```haskell
  main = print $ Main.reverse "Hello"  -- returns "olleH"
```

Remember, that strings are just lists of characters

- Implement a function which counts the number of odd numbers in a list of integers.
- Use pattern matching and recursion.

Type definition:
`countOdds :: [Int] -> Int`

Example main routine:
`main = print $ countOdds [1,2,3,4,5]  -- returns 3`

 
