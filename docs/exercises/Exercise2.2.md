# Exercise 2.2 - Basics

If you do not finish during the lecture period, please finish it as homework.

## Book Index (Working with maps, arrays and slices)

A Book Index is an inverted index which lists all pages a word occurs in a book.
Write a program which generates an inverted index out of an array of book pages.
Each Page contains an array of words.

- Use the "type" keyword to define custom types for Book, Page and Index. 
- Write a unit test which generates a book and calculates the index and prints it to Stdout

## After this Exercise
- You know the basic Go container types: string, map, array, slice
- You know how to make custom types printable

# Usage of the Library functions and error handling
Write a program "find" that searches the filesystem recursively from a given path and regex expression. 

1. Use the flag library (https://tour.golang.org/methods/17) to provide the following parameters 
```
   Usage of find:
   -path string
       path to search (default ".")
   -regex string
       path (default ".*")
```
2. Use the ioutil.ReadDir (https://pkg.go.dev/io/ioutil@go1.17.2#ReadDir) function to list the contents of a given directory. Check each file for the given regex with the 
   "regexp.MatchString" (https://pkg.go.dev/regexp#MatchString) function and print its path+name on the screen.
3. Either use "panic" or "log.Fatal" for error handling.
4. Run through directories recursively
 
# Question
Go doesn't support Exceptions but uses multiple return values of which one can be the error information.
Discuss the pro and cons about both approaches.
