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

### Part A - Yourself: Test and implementation

Do this part without a coding agent.

- Write a test first: sort ascending and descending with the same `BubbleSort` function.
- Then implement `BubbleSort` with the comparison function.
- The comparison function gets two *indices*, not two values. Where do the values come from? What is the name of this concept?

### Part B - With a coding agent: Generic version

Let a coding agent write a generic version:

```go
func BubbleSort[T any](data []T, less func(a, b T) bool)
```

Compare it with your version and with the standard library:

- `sort.Slice(x any, less func(i, j int) bool)` gets indices. `slices.SortFunc(x []E, cmp func(a, b E) int)` gets values. Why did `sort.Slice` have to use indices? Hint: `sort.Slice` is older than Go 1.18.
- Does your test from Part A still work with the generic version?

## Exercise 5.2a - Filter

Given a list of strings, write a function `Filter` which takes a 
filter function as argument and filters all elements of the array, which do 
not match the filter function.

What type must the filter function have?
Filter out all strings which contain digits.

E.g.

```go
    notDigit := func .....
    array := []string{"a", "b3b", "c", "12", "D"}
    result := Filter(array, notDigit)
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
    array := []string{"eVeRyThInG", "uPpErCaSe"}
    result := Map(array, toUppercase)
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
    array := []string{"a", "b", "c", "d"}
    result := Reduce(array, concat)
	fmt.Println(result) // "a,b,c,d" 
```

### Questions for 5.2

Answer these yourself, before you write or generate any code:

- Write down the function types for `Filter`, `Map` and `Reduce`.
- Is your `Map` pure? Does it change the input slice, or does it return a new one?
- What should `Reduce` return for an empty slice?
- Make the functions generic, e.g. `func Map[T, U any](data []T, f func(T) U) []U`. Predict: What does `Map([]string{"a", "bb", "ccc"}, func(s string) int { return len(s) })` return? Which type does it have?

## Exercise 5.3 - Map / Filter / Reduce

Expand the Exercise 5.2 to a Map / Filter / Reduce streaming solution.

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

### Generics

Let a coding agent convert your stream to generics (`Stream[T]`). Then check:

- How did the agent solve `Map` from `T` to `U`? Go methods cannot have their own type parameters. Is `Map` still a method?
- Which type assertions like `o.(string)` are gone?

### Is this a bug?

This is the `Map` method of the stream from the lecture (`src/functional/streams/streams.go`):

```go
func (s *SliceStream) Map(mapper Mapper) Stream {
    for i, e := range s.data {
        s.data[i] = mapper(e)
    }
    return s
}
```

Decide **without running it**: What does this print? Write down your answer and the reason.

```go
data := []any{"a", "b"}
result := ToStream(data).Map(toUpperCase).Reduce(concat)
fmt.Println(result, data)
```

- Then run it. Is this a bug? Is `Map` pure?
- Fix it and write a test that fails with the old version.
- Does your own stream from 5.3 have the same problem? Does the stream of the coding agent?
- Bonus: The stream computes every step immediately for all elements. How would you make it *lazy*, so that `Map` only runs for the elements that `Filter` keeps? Haskell, the next lecture, is lazy by default.

## Exercise 5.4 - Word Count (WC)

Word Count is a famous algorithm for demonstrating the power of distributed collections and functional programming. 
Word Count counts how often a word (string) occurs in a collection. It is easy to address that problem with shared state (a map), but
this solution does not scale well.
The question here is how to use a pure functional algorithm to enable parallel and distributed execution.

After running Word Count, you should get the following result:

```
INPUT:  []any{"a", "a", "b", "b", "D", "a"}
OUTPUT: "a:3, b:2, D:1, "
```

*Questions*
- How can you implement the problem with the already built Map()/Filter()/Reduce() functions?
- Write an Unit Test to prove that your solution works as expected!

### Part A - Yourself: Design

Do this part without a coding agent:

- What does `Map` produce for one word? What does the accumulator combine?
- Write the accumulator on paper. Is it associative?

### Part B - With a coding agent: Implementation and review

Let a coding agent implement Word Count with your stream API. Then review the result and the `sumInts` accumulator in `src/functional/streams/streams_test.go`:

- Is the accumulator pure, or does it change one of its arguments?
- Parallel execution reduces each part on its own and combines the partial results at the end. Test the accumulator with two partial results: `sumInts([]Pair{{"a", 1}, {"b", 1}}, []Pair{{"a", 1}, {"b", 1}})`. Predict the result first. Is it correct?
- Would the solution still be correct, if the input was split into two parts and reduced in parallel?

### Part C - Go 1.23

Since Go 1.23 the standard library has iterators (`iter.Seq`) and functions like `slices.Collect`, `slices.Sorted` and `maps.Keys`.

- Write Word Count with a plain `map[string]int` and print the words sorted with `slices.Sorted(maps.Keys(counts))`.
- Compare both solutions: Which one is easier to read? Which one can run in parallel?

## Exercise 5.5 - Lambda Calculus

```
TRUE  := λa.λb.a
FALSE := λa.λb.b
NOT   := λp.(p FALSE TRUE)
OR    := λp.λq.(p p q)
```

### Part A - Yourself: On paper

Do this part on paper, without a computer:

- Evaluate `NOT FALSE` step by step.
- Evaluate `OR FALSE TRUE` and `OR FALSE FALSE` step by step.
- Write `XOR` as a lambda expression. Hint: If `p` is TRUE, the result is `NOT q`. Otherwise it is `q`.

### Part B - In the browser

Open the JavaScript console of your browser (F12) and use the functions from the lecture:

```javascript
TRUE = a => b => a;
FALSE = a => b => b;
NOT = p => p(FALSE)(TRUE);
ToBool = f => f(true)(false);
```

- Write `OR` and your `XOR` from Part A in JavaScript.
- Print the truth tables for both with `ToBool` and check them against your results from Part A.
