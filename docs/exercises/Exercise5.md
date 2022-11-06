# Exercise 5 - Functional Programming in Go

If you do not finish during the lecture period, please finish it as homework.

## Exercise 5.1 - Warm Up

Write a Go program which shows the following concepts:

- Functions as Variables
- Anonymous Functions
- High Order Functions (functions as parameters or return values)
- Closures (https://en.wikipedia.org/wiki/Closure_(computer_programming))

## Exercise 5.2 - Filter

Given a list of strings, write a function `Filter` which takes a 
filter function as argument and filters all elements of the array, which do 
not match the filter function.

E.g.

```go
    array := []strings{"a", "b", "c", "1", "D"}
    result := Filter(strings, notDigit)
```

What type must the filter function have?
Filter out all strings which are not digits.
 
## Exercise 5.3 - Map / Filter / Reduce

Expand the Exercice 5.2 to a Map / Filter / Reduce streaming solution.

Map/Reduce is a famous functional construct implemented in many parallel and distributed collection frameworks like
Hadoop, Apache Spark, Java Streams (not distributed but parallel), C# Linq

- Implement a custom M/R API with the following interface:
 ```go
    type Stream interface {
    	Map(m Mapper) Stream
    	Filter(p Predicate) Stream
    	Reduce(a Accumulator) any
    }
```

The usage of the interface should be like this:
```go
    stringSlice := []any{"a", "b", "c", "1", "D"}

	// Map/Reduce
	result := ToStream(stringSlice).
		Map(toUpperCase).
		Filter(notDigit).
		Reduce(concat).(string)

	if result != "A,B,C,D" {
		t.Error(fmt.Sprintf("Result should be 'A,B,C,D' but is: %v", result))
    }
```

 *Questions*
 - What is the type of Mapper, Predicate and Accumulator?
 - How can you make the types generic, so they work for any type, not only for string?

## Exercise 5.4 - Word Count (WC)

Word Count is a famous algorithm for demonstrating the power of distributed collections and functional programming. 
Word Count counts how often a word (string) occurs in a collection. It is easy to address that problem with shared state (a map), but
this solution does not scale well.
The question here is how to use a pure functional algorithm to enable parallel and distributed execution.

After running Word Count, you should get the following result:

INPUT:  "A" "B" "C" "A" "B" "A"

OUTPUT: ("A":3) ("B":2) ("C":1) 

*Questions*
- How can you implement the problem with the already built Map()/Filter()/Reduce() functions?
- Write an Unit Test to prove that your solution works as expected!
