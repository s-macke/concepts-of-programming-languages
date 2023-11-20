# Exercise 7 - Concurrent Programming in Go

If you do not finish during the lecture period, please finish it as homework.

## Exercise 7.1 - Waitgroup

The following code starts 10000 threads that increment a counter.

```go
var myYoutubevideo struct {
    likes int32
}

func Viewer() {
    time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
    atomic.AddInt32(&myYoutubevideo.likes, 1)
}

func main() {
    for i := 0; i < 10000; i++ {
        go Viewer()
    }
    for i := 0; i < 8; i++ {
        fmt.Println(myYoutubevideo.likes, "likes")
        time.Sleep(200 * time.Millisecond)
    }
}
```

Use a waitgroup to wait for all threads to finish before printing the final result.

## Exercise 7.2 - Generator

Write a generator for Fibonacci numbers, i.e. a function that returns a channel where the next Fibonacci number can be read.
```go
func main() {
    fibChan := fib() // <- write func fib
    for n := 1; n <= 10; n++ {
        fmt.Printf("The %dth Fibonacci number is %d\n", n, <-fibChan)
    }
}
```

## Exercise 7.3 - Timeout

Write a function `setTimeout()` that times out an operation after a given amount of time. Hint: Have a look at the built-in `time.After()` function and make use of the `select` statement.
```go
func main() {
    res, err := setTimeout(func() int {
        time.Sleep(2000 * time.Millisecond)
        return 1
    }, 1*time.Second)

    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Printf("operation returned %d", res)
    }
}
```

Hint: Within the function setTimeout, use two channels and wait for both channels with a select to return a value.
One channel is the result of the function, the other channel is a timer channel that is created with time.After().

## Exercise 7.4 - Dining Philosophers

Write a program to simulate the [Dining Philosophers](https://en.wikipedia.org/wiki/Dining_philosophers_problem) 
Problem by using Go Channels.
- There should be one Go Routine for each Philosopher
- The table itself should be a single Go Routine and should handle all forks via channels and a single select statement. This makes synchronization easier.
- Make sure that:
  - The distribution of forks is fair - No philosopher dies on starvation 
  - Use the given Unit Test:

```go
func TestPhilosophers(t *testing.T) {

	var COUNT = 5

	// start table for 5 philosophers
	table := NewTable(COUNT)

	// create 5 philosophers and run parallel 
	for i := 0; i < COUNT; i++ {
		philosopher := NewPhilosopher(i, table)
		go philosopher.run()
	}
	go table.run()

	// simulate 10 milliseconds --> check output
	time.Sleep(10 * time.Millisecond)
}
```

Sample console output:

```sh
[->]: Philosopher #0 eats ...
[->]: Philosopher #3 eats ...
[<-]: Philosopher #0  eat ends.
[<-]: Philosopher #3  eat ends.
[->]: Philosopher #0 thinks ...
[->]: Philosopher #1 eats ...
[->]: Philosopher #3 thinks ...
[->]: Philosopher #4 eats ...
[<-]: Philosopher #1  eat ends.
[->]: Philosopher #1 thinks ...
[<-]: Philosopher #4  eat ends.
[->]: Philosopher #2 eats ...
[->]: Philosopher #4 thinks ...
[<-]: Philosopher #0 thinking ends
[->]: Philosopher #0 eats ...
[<-]: Philosopher #3 thinking ends
```
