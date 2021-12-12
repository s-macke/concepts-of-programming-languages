package main

import (
	"time"
)

func main() {

	var COUNT = 5

	// start table for 5
	table := NewTable(COUNT)

	// start philosophers
	for i := 0; i < COUNT; i++ {
		philosopher := NewPhilosopher(i, table)
		go philosopher.run()
	}
	go table.run()

	// wait 1 millisecond --> check output
	time.Sleep(10 * time.Second)
}
