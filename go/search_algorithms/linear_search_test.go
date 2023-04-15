package searchalgorithms

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

// LinearSearchData stores data needed to test the “LinearSearch“ function.
type LinearSearchData[T comparable] struct {
	slice  []T
	value  T
	result int
}

// TestLinearSearch_Int tests the function `LinearSearch` with integer input.
func TestLinearSearch_Int(t *testing.T) {
	// Get a new assertion object for use within this function
	assert := assert.New(t)

	// Get a slice of structs defining the test data
	data := []LinearSearchData[int]{
		// Test for value data that occurs once
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 5, 5},
		// Test for value data that occurs twice
		{[]int{0, 1, 1, 5, 5, 6, 6, 7, 8, 9}, 5, 3},
		// Test for empty slice
		{[]int{}, 0, NotFound},
		// Test for nil slice
		{nil, 10, NotFound},
	}

	for _, datum := range data {
		reval, err := LinearSearch(datum.slice, datum.value)
		if err != nil {
			log.Panic(err)
		}
		assert.Equal(datum.result, reval)
	}
}

// TestLinearSearch_String tests the function `LinearSearch` with string input.
func TestLinearSearch_String(t *testing.T) {
	// Get a new assertion object for use within this function
	assert := assert.New(t)

	data := []LinearSearchData[string]{
		// Test for data which occurs once
		{[]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}, "d", 3},
		// Test for data longer that one byte
		{[]string{"cheezaram", "okeke"}, "okeke", 1},
		// Test for data occuring more than once
		{[]string{"okeke", "okeke", "cheezaram", "cheezaram"}, "cheezaram", 2},
		// Test empty slice
		{[]string{}, "cheezaram", NotFound},
		// Test nil slice
		{nil, "okeke", NotFound},
	}

	for _, datum := range data {
		reval, err := LinearSearch(datum.slice, datum.value)
		if err != nil {
			log.Panic(err)
		}
		assert.Equal(datum.result, reval)
	}
}
