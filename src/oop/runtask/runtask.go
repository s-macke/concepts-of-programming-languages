package main

import (
	"fmt"
)

type Task struct {
	id int
}

func (t Task) Run() {
	fmt.Println("Running task", t.id)
}

// Runner is the "base class"
type Runner struct{}

func (r *Runner) Run(task Task) {
	fmt.Println("Run task id", task)
}

func (r *Runner) RunAll(tasks []Task) {
	for _, t := range tasks {
		r.Run(t)
	}
}

// RunCounter "extends" Runner via embedding
type RunCounter struct {
	Runner
	count int
}

func (rc *RunCounter) Run(task Task) {
	rc.count++
	rc.Runner.Run(task)
}

func (rc *RunCounter) RunAll(tasks []Task) {
	// batch increment like the Java override
	rc.count += len(tasks)
	rc.Runner.RunAll(tasks)
}

func (rc *RunCounter) Count() int { return rc.count }

func main() {
	tasks := []Task{{id: 1}, {id: 2}, {id: 3}}

	r := &Runner{}
	r.RunAll(tasks)
	/*
		rc := &RunCounter{}
		rc.RunAll(tasks)

		fmt.Printf("Tasks executed: %d\n", rc.Count())
	*/
}
