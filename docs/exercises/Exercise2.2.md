# Exercise 2.2 - Basics

If you do not finish during the lecture period, please finish it as homework.

## Book Index (Working with maps, arrays and slices)

A Book Index is an inverted index which lists all pages a word occurs in a book.
Write a program which generates an inverted index out of an array of book pages.
Each Page contains an array of words.

- Use the "type" keyword to define custom types for Book, Page and Index. 
- Write a unit test which generates a book and calculates the index and prints it to Stdout

### Part A - Yourself: Types and test

Do this part without a coding agent.

- Define the types `Book`, `Page` and `Index` yourself.
- Write a test with a tiny book of three pages. Write the **expected** index yourself as a map literal and compare it with the calculated index via `reflect.DeepEqual`.
- Why can't you compare the two indexes with `==`?

### Part B - With a coding agent: Implementation

Let a coding agent implement the function that calculates the index. Your test decides whether the result is correct. Then check:

- If a word occurs twice on the same page, is the page listed twice in the index? Does your test check this case at all?
- Are the page numbers sorted?
- Iterate over the index with `range` and print every entry. Run the program several times. Why does the order change? Why did the Go designers do this on purpose?

### Part C - Comparison with Java

Let the coding agent write the same index in Java and compare both versions.

- Java needs something like `computeIfAbsent` or `containsKey`. Why is `index[word] = append(index[word], page)` enough in Go?
- Hint: What does a map return for a missing key? And what does `append` do with a nil slice?

## After this Exercise

- You know the basic Go container types: string, map, array, slice
