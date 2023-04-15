package searchalgorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

// JumpSearchData is a struct for storing the data for use “JumpSearch“
type JumpSearchData[T constraints.Ordered] struct {
	slice  []T
	value  T
	result int
}

// TestJumpSearch_Int tests the `JumpSearch` function for integer values
func TestJumpSearch_Int(t *testing.T) {
	assert := assert.New(t)

	data := []JumpSearchData[int]{
		// Test value that occurs once
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 3, 3},
		// Test value that occurs multiple times
		{[]int{0, 1, 2, 2, 3, 3, 5, 7, 8, 11, 12, 12, 12, 23}, 12, 10},
		// Test emtpy slice
		{[]int{}, 0, NotFound},
		// Test nil slice
		{nil, 0, NotFound},
	}

	for _, datum := range data {
		reval, err := JumpSearch(datum.slice, datum.value)
		if assert.NoError(err) {
			assert.Equal(datum.result, reval)
		}
	}
}

// TestJumpSearch_String tests the `JumpSearch` function for string values
func TestJumpSearch_String(t *testing.T) {
	assert := assert.New(t)

	data := []JumpSearchData[string]{
		// Test value that occurs once
		{[]string{"a", "b", "c", "d", "e", "f", "g", "h"}, "b", 1},
		// Test value that occurs multiple times
		{[]string{"cheezaram", "cheezaram", "okeke"}, "cheezaram", 0},
		// Test slice with no matching value
		{[]string{"emmanuel", "cheezaram", "okeke"}, "ovy", NotFound},
		// Test empty slice
		{[]string{}, "", NotFound},
		// Test nil slice
		{nil, "", NotFound},
	}

	for _, datum := range data {
		reval, err := JumpSearch(datum.slice, datum.value)
		if assert.NoError(err) {
			assert.Equal(datum.result, reval)
		}
	}
}

func TestJumpSearch_Float(t *testing.T) {
	assert := assert.New(t)

	data := []JumpSearchData[float64]{
		// Test for value data that occurs once
		{[]float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9, 10.1}, 9.9, 8},
		// Test for value data that occurs twice
		{[]float64{0.1, 1.2, 2.3, 3.4, 4.5, 5.6, 5.6, 6.7, 7.8, 8.9}, 5.6, 5},
		// Test for value data that does not exist in slice
		{[]float64{0.1, 1.2, 2.3, 3.4, 4.5, 5.6, 6.7, 7.8, 8.9}, 10.0, NotFound},
		// Test for empty slice
		{[]float64{}, 0.0, NotFound},
		// Test for nil slice
		{nil, 10.0, NotFound},
	}

	for _, datum := range data {
		reval, err := JumpSearch(datum.slice, datum.value)
		if assert.NoError(err) {
			assert.Equal(datum.result, reval)
		}
	}
}
