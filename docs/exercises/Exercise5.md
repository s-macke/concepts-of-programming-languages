# Exercise 5 - Functional Programming in Go

If you do not finish during the lecture period, please finish it as homework.

## Exercise 5.1 - Warm Up

This is a Bubble Sort algorithm for ints.

```go
package main

import "fmt"

// Sorts an array of integers using the bubble sort algorithm.
func BubbleSort(data []int) {
    for i := 0; i < len(data); i++ {
        for j := 0; j < len(data)-1; j++ {
            if data[j] > data[j+1] {
                data[j], data[j+1] = data[j+1], data[j] // Swap the values
            }
        }
    }
}

func main() {
    data := []int{27, 15, 8, 9, 12, 4, 17, 19, 21, 23, 25}
    BubbleSort(data)
    fmt.Println(data)
}
```

Rewrite it, so that the Bubble Sort algorithm takes a comparison function as input
```go
BubbleSort(data, func(i, j int) bool {return data[i] > data[j]})
```


## Exercise 5.2a - Filter

Given a list of strings, write a function `Filter` which takes a 
filter function as argument and filters all elements of the array, which do 
not match the filter function.

What type must the filter function have?
Filter out all strings which contain digits.

E.g.

```go
    notDigit := func .....
    array := []strings{"a", "b3b", "c", "12", "D"}
    result := Filter(strings, notDigit)
	fmt.Println(result) // "a", "c", "D" 
```

## Exercise 5.2b - Map

Given a list of strings, write a function `Map` which takes a
`map` function as argument. The `Map` function applies the `map` function to all elements of the array. 

What type must the `map` function have?

Write the `map` function to convert a string to uppercase.
E.g.

```go
    toUppercase := func .....
    array := []strings{"eVeRyThInG", "uPpErCaSe"}
    result := Map(strings, toUppercase)
    fmt.Println(result) // "EVERYTHING", "UPPERCASE" 
```

## Exercise 5.2c - Reduce

Given a list of strings, write a function `Reduce` which takes a
`reduce` function as argument. `Reduce` accumulates all elements of the 
array into a single value.
The `reduce` function takes two arguments, the first is the current result, 
the second is the current element of the array.

```go
    concat := func concat(a string, b string) string {
        return a + "," + b
    }
    array := []strings{"a", "b", "c", "d"}
    result := Reduce(strings, concat)
	fmt.Println(result) // "a,b,c,d" 
```

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
Hint: If you are unfamiliar with the type `any`, limit the Stream 
elements to a single type like a `string`.

The usage of the interface should be like this:
```go
    stringSlice := []any{"a", "b", "c", "1", "D"}

	// Map/Reduce
	result := ToStream(stringSlice).
		Map(toUpperCase).
		Filter(notDigit).
		Reduce(concat)

	if result != "A,B,C,D" {
		t.Error(fmt.Sprintf("Result should be 'A,B,C,D' but is: %v", result))
    }
```

*Question*
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
