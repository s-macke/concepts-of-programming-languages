// Copyright 2018 Johannes Weigend
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package philosophers implements the Dining Philosophers Problem.
package main

// =====================================================================================================================
// Dining Philosophers Problem. See https://en.wikipedia.org/wiki/Dining_philosophers_problem.
// The synchronization is done with channels. There are two channels for put and request fork operation.
// There is one channel per philosopher to signal when the forks can be taken.
// =====================================================================================================================

// Table represents the table with dynamic seat count.
type Table struct {
	// channel for take requests - answer is sent over the reservedCh.
	takeCh chan int
	// channel to put fork requests.
	putCh chan int
	// channel to get signal if forks are available
	reservedCh []chan bool
	forkInUse  []bool
	nbrOfSeats int
}

// NewTable constructs a table with n seats.
func NewTable(nbrOfSeats int) *Table {
	table := &Table{
		nbrOfSeats: nbrOfSeats,
		forkInUse:  make([]bool, nbrOfSeats),
		takeCh:     make(chan int),
		putCh:      make(chan int),
	}

	table.reservedCh = make([]chan bool, nbrOfSeats)
	for i := 0; i < nbrOfSeats; i++ {
		table.reservedCh[i] = make(chan bool)
	}
	return table
}

func (t *Table) askForForks(id int) bool {
	t.takeCh <- id
	return <-t.reservedCh[id]
}

func (t *Table) putForks(id int) {
	t.putCh <- id
}

func (t *Table) SetForksInUse(id int, inuse bool) {
	t.forkInUse[id] = inuse
	t.forkInUse[(id+1)%t.nbrOfSeats] = inuse
}

func (t *Table) AreForksAvailable(id int) bool {
	return !t.forkInUse[id] && !t.forkInUse[(id+1)%t.nbrOfSeats]
}

// Function run() contains the main loop for assigning forks and starting philosophers.
func (t *Table) run() {
	for {
		select {
		case requestedForksId := <-t.takeCh:
			{
				if t.AreForksAvailable(requestedForksId) {
					// both forks are not in use -> reserve
					t.SetForksInUse(requestedForksId, true)
					t.reservedCh[requestedForksId] <- true
				} else {
					t.reservedCh[requestedForksId] <- false // not valid try again --> see loop in takeForks.
				}
			}
		case putForksId := <-t.putCh:
			{
				// put forks
				t.SetForksInUse(putForksId, false)
			}
		}
	}
}
