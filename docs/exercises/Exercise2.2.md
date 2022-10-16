# Exercise 2.2 - Basics

If you do not finish during the lecture period, please finish it as homework.

## Book Index (Working with maps, arrays and slices)

A Book Index is an inverted index which lists all pages a word occurs in a book.
Write a program which generates an inverted index out of an array of book pages.
Each Page contains an array of words.

- Define custom types for Book, Page and Index
- Make sure the Stringer() interface is implemented for Book and Index to make them printable
  More details about the Stringer interface: https://tour.golang.org/methods/17
  The stringer interface will be explained in more detail in the next lecture.
- Write a unit test which generates a book and calculates the index and prints it to Stdout

## After this Exercise
- You know the basic Go container types: string, map, array, slice
- You know how to make custom types printable

# Usage of the Library functions and error handling
Write a program "find" that searches the filesystem recursively from a given path and regex expression. 

1. Use the flag library to provide the following parameters 
```
   Usage of find:
   -path string
       path to search (default ".")
   -regex string
       path (default ".*")
```

3. Use the ioutil.ReadDir (https://pkg.go.dev/io/ioutil@go1.17.2#ReadDir) function to list the contents of a given directory. Check each file for the given regex with the 
   "regexp.MatchString" (https://pkg.go.dev/regexp#MatchString) function and print its path+name on the screen.
4. Either use "panic" or "log.Fatal" for error handling.
5. Run through directories recursively
 
# Question
Go doesn't support Exceptions but uses multiple return values of which one can be the error information.
Discuss the pro and cons about both approaches.
