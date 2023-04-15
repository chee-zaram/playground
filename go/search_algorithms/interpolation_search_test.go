package searchalgorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// InterpolationSearchData stores data needed to perform tests.
type InterpolationSearchData struct {
	slice  []int
	value  int
	result int
}

// TestInterpolationSearch performs tests for InterpolationSearch.
func TestInterpolationSearch(t *testing.T) {
	// Get new assertion object for use in testing
	assert := assert.New(t)

	data := []InterpolationSearchData{
		// Test with value that occurs once
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		// Test with value that occurs multiple times
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 8, 8, 9}, 8, 8},
		// Test with value that does not exist in slice
		{[]int{2, 3, 4, 5, 5, 5, 6, 7}, 23, NotFound},
		// Test with empty slice
		{[]int{}, 0, NotFound},
		// Test with nil slice
		{nil, 0, NotFound},
	}

	for _, datum := range data {
		reval, err := InterpolationSearch(datum.slice, datum.value)
		if assert.NoError(err) {
			assert.Equal(datum.result, reval)
		}
	}
}
