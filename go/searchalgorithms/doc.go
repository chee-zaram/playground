// Package searchalgorithms implements many of the most used search algorithms
// using the Go programming language.
//
// It utilizes the latest and most efficient techniques, like generics, in order
// to make functions work for as many types as possible.
//
// Functions typically take in a slice of values and a value to search for.
// They typically return the index of the (first occurrence, where applicable, of the)
// value and an error (if any) that occurred.
//
// `NB`: An error is not returned if value is not found, rather the index returned
// is `NotFound` which is -1. An error is only returned when the log file could not be
// created, opened, or written to.
package searchalgorithms
