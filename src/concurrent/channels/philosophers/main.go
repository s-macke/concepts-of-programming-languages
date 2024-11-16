package main

func main() {
	// start table for 5
	nbrOfSeats := 5
	table := NewTable(nbrOfSeats)

	// start philosophers
	for i := 0; i < nbrOfSeats; i++ {
		philosopher := NewPhilosopher(i, table)
		go philosopher.run()
	}
	table.run()
}
