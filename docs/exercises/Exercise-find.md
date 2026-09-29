# Exercise - find (parked)

Moved out of Exercise 2.2. To be used in a later lecture about error handling.

## Usage of the Library Functions and Error Handling

Write a program "find" that searches the file system recursively, starting from a given path, for file names matching a regular expression.

1. Use the flag library (<https://pkg.go.dev/flag>) to provide the following parameters:
```
   Usage of find:
   -path string
       path to search (default ".")
   -regex string
       regular expression for file names (default ".*")
```
2. Use the function `os.ReadDir` (<https://pkg.go.dev/os#ReadDir>) to list the contents of a directory. Check each file name against the given regex with `regexp.MatchString` (<https://pkg.go.dev/regexp#MatchString>) and print its path and name.
3. Use either `panic` or `log.Fatal` for error handling.
4. Search the directories recursively. Implement the recursion yourself, do not use `filepath.WalkDir`.

## Question

Go doesn't support exceptions, but uses multiple return values, one of which can be the error information.
Discuss the pros and cons of both approaches.

## After this Exercise

- You know how to use library functions such as `flag`, `os` and `regexp`
- You know how errors are handled in Go
