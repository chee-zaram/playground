package searchalgorithms

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

// BinarySearchData stores data needed to test the `BinarySearch` function.
type BinarySearchData[T constraints.Ordered] struct {
	slice  []T
	value  T
	result []int
}

// TestBinarySearch_Int tests the “BinarySearch“ function with integer slices.
func TestBinarySearch_Int(t *testing.T) {
	// Get a new assertion object for use throughout the function
	assert := assert.New(t)

	// Define a list of struct data
	data := []BinarySearchData[int]{
		// Test data that occurs once in slice
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10, []int{9}},
		// Test data that occurs twice in slice
		{[]int{1, 2, 3, 4, 5, 5, 7, 8, 9, 10}, 5, []int{4, 5}},
		// Test empty slice
		{[]int{}, 10, []int{NotFound}},
		// Test nil slice
		{nil, 0, []int{NotFound}},
	}

	for _, datum := range data {
		reval, err := BinarySearch(datum.slice, datum.value)
		if err != nil {
			log.Panic(err)
		}
		assert.Contains(datum.result, reval)
	}
}

// TestBinarySearch_String tests the “BinarySearch“ function with string slices.
func TestBinarySearch_String(t *testing.T) {
	assert := assert.New(t)

	data := []BinarySearchData[string]{
		// Test data that occurs once in slice
		{[]string{"a", "b", "c", "d", "e", "f", "g"}, "e", []int{4}},
		// Test data that occurs twice in slice
		{[]string{"a", "b", "c", "d", "d", "d", "e", "f"}, "d", []int{3, 4, 5}},
		// Test empty slice
		{[]string{}, "a", []int{NotFound}},
		// Test nil slice
		{nil, "", []int{NotFound}},
	}

	for _, datum := range data {
		reval, err := BinarySearch(datum.slice, datum.value)
		if err != nil {
			log.Panic(err)
		}
		assert.Contains(datum.result, reval)
	}
}

// TestBinarySearch_Float tests the “BinarySearch“ function with float slices.
func TestBinarySearch_Float(t *testing.T) {
	assert := assert.New(t)

	data := []BinarySearchData[float64]{
		{[]float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.1}, 9.9, []int{8}},
		{[]float64{2.0, 2.2, 4.4, 4.4, 5.0, 6.6}, 4.4, []int{2, 3}},
	}

	for _, datum := range data {
		reval, err := BinarySearch(datum.slice, datum.value)
		if err != nil {
			log.Panic(err)
		}
		assert.Contains(datum.result, reval)
	}
}
