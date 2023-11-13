# Exercise 6 - Haskell

Use an online Haskell interpreter such as.

* https://replit.com/languages/haskell (Login required)
* https://www.jdoodle.com/execute-haskell-online/
* https://www.onlinegdb.com/online_haskell_compiler

Solve the following exercises.

### Exercise 6.1 - Guards
- Write A signum function which returns 1 if the number is positive, -1 if the number is negative and 0 if the number is 0. 
- Write a function with guards that returns the signum of a number.

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
  Use pattern matching and recursion.

Type definition:

`countOdds :: [Int] -> Int`

Example main routine:

`main = print $ countOdds [1,2,3,4,5]  -- returns 3`

### Exercise 6.4 - Higher Order Functions

- Implement your own filter function which takes a list and a 
  predicate function and returns a list of all elements that satisfy the predicate.
  For example, the expression

```haskell
   main = print $ filter even [1,2,3,4,5] 
```
returns [2,4]

### Exercise 6.5 - Random number guessing game

This game generates a random number between 1 and 100, and the player has to guess it. 
After each guess, the game will tell the player if their guess is too high, too low, or correct. 
The game continues until the player guesses the correct number.

Implement this game in Haskell.

#### Excercise 6.5a - random numbers

Check that the function getRandomFromTime indeed returns a random number.

```haskell
import Data.Time.Clock
import Data.Time.Clock.POSIX

-- returns a random number between 1 and 100
getRandomFromTime :: IO Int
getRandomFromTime = do
    now <- getCurrentTime
    let millis = floor $ utcTimeToPOSIXSeconds now * 1000
    return $ (millis `mod` 100) + 1

main :: IO ()
main = do secret <- getRandomFromTime
          print secret

```

#### Excercise 6.6b - Determine status

Given the main function 

```haskell
main :: IO ()
main = do secret <- getRandomFromTime
          putStrLn "Try to guess a secret number 1-100."
          input <- getLine
          let guess = read input :: Int
          isWin <- Main.status guess secret
          print isWin
          putStrLn "Done"
```

Implement the status function, which returns a boolean if the guess is correct.
It should also print a message if the guess is too high or too low or it you have won
The function should have the following signature:

```haskell
status :: Int -> Int -> IO Bool
```

#### Excercise 6.6c - Loop

Finally implement a loop until the user has guessed the correct number.