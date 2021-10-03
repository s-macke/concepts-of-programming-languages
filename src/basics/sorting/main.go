package main

import (
	"log"
	"sort"
)

// AxisSorter sorts planets by axis.
type AxisSorter []Planet

func (a AxisSorter) Len() int           { return len(a) }
func (a AxisSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a AxisSorter) Less(i, j int) bool { return a[i].Axis < a[j].Axis }

// NameSorter sorts planets by name.
type NameSorter []Planet

func (a NameSorter) Len() int           { return len(a) }
func (a NameSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a NameSorter) Less(i, j int) bool { return a[i].Name < a[j].Name }

// Planet represents a planet in our solar system.
type Planet struct {
	Name       string
	Aphelion   float64 // in million km
	Perihelion float64 // in million km
	Axis       int64   // in km
	Radius     float64
}

func main() {
	mars := Planet{
		Name:       "Mars",
		Aphelion:   249.2,
		Perihelion: 206.7,
		Axis:       227939100,
		Radius:     3389.5,
	}

	earth := Planet{
		Name:       "Earth",
		Aphelion:   151.930,
		Perihelion: 147.095,
		Axis:       149598261,
		Radius:     6371.0,
	}

	venus := Planet{
		Name:       "Venus",
		Aphelion:   108.939,
		Perihelion: 107.477,
		Axis:       108208000,
		Radius:     6051.8,
	}

	planets := []Planet{mars, venus, earth}
	log.Println("unsorted:", planets)

	sort.Sort(AxisSorter(planets))
	log.Println("by axis:", planets)

	sort.Sort(NameSorter(planets))
	log.Println("by name:", planets)
}
