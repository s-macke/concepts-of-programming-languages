// Copyright 2018 Johannes Weigend
// Licensed under the Apache License, Version 2.0

// Package streams contains a minimal Java8 Streams like API for slices.
// Usage:
// result := ToStream(slice).
//
//			Map(toUpperCase).
//			Filter(containsDigit).
//			Reduce(concat).(string)
//	 - ToStream() converts a Slice of "any" to a Steam.
//	 - Map() applies a function to all elements.
//	 - Filter() filters all elements out which not match the given predicate.
//	 - Reduce() combines all elements and returns a single element.
//
// This is a very first draft with no lazy or parallel support.
package streams

// EXERCISE 5.3 START OMIT

// Predicate function returns true if a given element should be filtered.
type Predicate func(any) bool

// Mapper function maps a value to another value.
type Mapper func(o1 any) any

// Accumulator function returns a combined element.
type Accumulator func(any, any) any

// EXERCISE 5.3 END OMIT

// Pair of two values.
type Pair struct {
	k any
	v any
}

// Stream interface is implemented for container types.
type Stream interface {
	Map(m Mapper) Stream
	Filter(p Predicate) Stream
	Reduce(a Accumulator) any
}

// ToStream helper converts a slice into a Stream.
func ToStream(s []any) Stream {
	return NewSliceStream(s)
}

// SliceStream is a stream implementation for slices.
type SliceStream struct {
	data []any
}

// NewSliceStream returns a new stream.
func NewSliceStream(data []any) *SliceStream {
	ss := &SliceStream{}
	ss.data = data
	return ss
}

// Map applies the Mapper on all elements.
func (s *SliceStream) Map(mapper Mapper) Stream {
	for i, e := range s.data {
		s.data[i] = mapper(e)
	}
	return s
}

// Filter filters all elements out.
func (s *SliceStream) Filter(filter Predicate) Stream {
	data := new([]any)
	for _, e := range s.data {
		if filter(e) {
			*data = append(*data, e)
		}
	}
	s.data = *data
	return s
}

// Reduce combines two elements and return one element.
func (s *SliceStream) Reduce(accumulate Accumulator) any {

	var result any
	for i, e := range s.data {
		if i == 0 {
			result = e
		} else {
			result = accumulate(result, s.data[i])
		}
	}
	return result
}
