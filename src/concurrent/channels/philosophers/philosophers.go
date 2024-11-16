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

package main

import (
	"fmt"
	"time"
)

// Philosopher represents one philosopher.
type Philosopher struct {
	id    int
	table *Table
}

// NewPhilosopher constructs a philosopher.
func NewPhilosopher(id int, table *Table) *Philosopher {
	return &Philosopher{
		id:    id,
		table: table,
	}
}

// Run loops forever.
func (p *Philosopher) run() {
	for {
		p.takeForks()
		p.eat()
		p.putForks()
		p.think()
	}
}

// Take forks by channeling our id to the table and wait until the table returns true on the reserved channel.
func (p *Philosopher) takeForks() {
	fmt.Printf("Philosopher #%d is getting hungry\n", p.id)
	// try to get forks from table
	gotForks := false
	for !gotForks {
		gotForks = p.table.askForForks(p.id)
	}
}

// Put forks by channeling our id to the table. The table is responsible for the put logic.
func (p *Philosopher) putForks() {
	fmt.Printf("Philosopher #%d puts down forks\n", p.id)
	p.table.putForks(p.id)
}

// Eating.
func (p *Philosopher) eat() {
	fmt.Printf("Philosopher #%d starts eating\n", p.id)
	time.Sleep(2 * time.Second)
}

// Thinking.
func (p *Philosopher) think() {
	fmt.Printf("Philosopher #%d is thinking\n", p.id)
	time.Sleep(3 * time.Second)
}
