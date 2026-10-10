# Exercise 2.1 - Basics

If you do not finish during the lecture period, please finish it as homework.

## Pointers (swap.go)

Start with the test (last bullet) before you implement the functions. The test defines what "swapped" means: Are the values swapped, or the pointers?

- Write a function to swap two integer variables in the caller function
- Write a function to swap two integer pointers in the caller function, so that each pointer points to the other variable afterwards. Find two ways: returning the swapped pointers, and passing pointers to the pointers.
- Write a unit test `TestSwap` in a file `swap_test.go` and run it with `go test`

## Is this a bug?

Each of the following functions looks plausible. For each one:

1. Decide **without running it**: Is it a bug? Write down your answer and the reason.
2. Run it and check your answer.
3. Ask a coding agent the same question and compare its explanation with yours.

```go
// Should set the caller's pointer to nil.
func reset(p *int) {
    p = nil
}
```

```go
type Item struct {
    Name  string
    Price int
}

// Should halve the price of every item.
func discount(items []Item) {
    for _, it := range items {
        it.Price = it.Price / 2
    }
}
```

```go
// Should return a new counter that starts at 0.
func newCounter() *int {
    x := 0
    return &x
}
```

Hint: The last one would be a classic bug in C. Is it a bug in Go?

### Questions

- What is the difference between a pointer and a Java object reference?
- What happens if you assign nil to a pointer variable and dereference the pointer?
- Did the coding agent explain all three functions correctly? Which concept from the lecture did you need to judge its answer?

## After this Exercise

- You know the difference between pointers and values
- You know how to write basic unit tests
