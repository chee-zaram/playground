package searchalgorithms

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

// ExponentialSearchData stores the needed to test “ExponentialSearch“
type ExponentialSearchData[T constraints.Ordered] struct {
	// slice containing values to check in
	slice []T
	// value to search for
	value T
	// the expected index of the value
	result int
}

// TestExponentialSearch_Int tests the “ExponentialSearch“ algorithm with int type.
func TestExponentialSearch_Int(t *testing.T) {
	assert := assert.New(t)

	data := []ExponentialSearchData[int]{
		// Test with target that occurs once
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 9, 9},
		// Test with target that occurs multiple times
		{[]int{0, 1, 2, 3, 4, 4, 5, 6, 8, 10, 12, 12, 12, 23, 34, 34}, 12, 10},
		// Test with target not in slice
		{[]int{13, 14, 15, 19}, 12, NotFound},
		// Test int32 lower limit
		{[]int{math.MinInt32, -1, 0, 1, math.MaxInt32}, math.MinInt32, 0},
		// Test int32 upper limit
		{[]int{math.MinInt32, -1, 0, 1, math.MaxInt32}, math.MaxInt32, 4},
		// Test int64 lower limit
		{[]int{math.MinInt64, -1, 0, 1, math.MaxInt64}, math.MinInt64, 0},
		// Test int64 upper limit
		{[]int{math.MinInt64, -1, 0, 1, math.MaxInt64}, math.MaxInt64, 4},
		// Test with empty slice
		{[]int{}, 0, NotFound},
		// Test with nil slice
		{nil, 0, NotFound},
	}

	for _, datum := range data {
		reval, err := ExponentialSearch(datum.slice, datum.value)
		if assert.NoError(err) {
			assert.Equal(datum.result, reval)
		}
	}
}

// TestExponentialSearch_Float tests the “ExponentialSearch“ algorithm with float64 type.
func TestExponentialSearch_Float(t *testing.T) {
	assert := assert.New(t)

	data := []ExponentialSearchData[float64]{
		// Test with target that occurs once
		{[]float64{0.1, 1.2, 2.3, 3.4, 4.5, 5.6}, 2.3, 2},
		// Test with target that occurs multiple times
		{[]float64{0.1, 1.2, 2.3, 3.4, 4.5, 4.5, 4.5, 5.6}, 4.5, 4},
		// Test with target that is not in slice
		{[]float64{0.1, 1.2, 2.3, 3.4, 4.5, 4.5, 4.5, 5.6}, 2.4, NotFound},
		// Test upper limit
		{[]float64{math.SmallestNonzeroFloat64, 0.1, 1.2, math.MaxFloat64}, math.MaxFloat64, 3},
		// Test lower limit
		{[]float64{math.SmallestNonzeroFloat64, 0.1, 1.2, math.MaxFloat64}, math.SmallestNonzeroFloat64, 0},
		// Test with empty slice
		{[]float64{}, 0.1, NotFound},
		// Test with nil slice
		{nil, 0.0, NotFound},
	}

	for _, datum := range data {
		reval, err := ExponentialSearch(datum.slice, datum.value)
		if assert.NoError(err) {
			assert.Equal(datum.result, reval)
		}
	}
}

// TestExponentialSearch_String tests the “ExponentialSearch“ algorithm with string type.
func TestExponentialSearch_String(t *testing.T) {
	assert := assert.New(t)

	data := []ExponentialSearchData[string]{
		// Test with target that occurs once
		{[]string{"a", "b", "c", "d", "e", "f", "g"}, "e", 4},
		// Test with target that occurs multiple times
		{[]string{"a", "b", "c", "d", "e", "f", "g", "g", "g", "g", "h", "i", "j"}, "g", 6},
		// Test with target of more than one byte
		{[]string{"cheezaram", "emmanuel", "emmanuel", "okeke"}, "emmanuel", 1},
		// Test with target not in slice
		{[]string{"cheezaram", "emmanuel", "okeke"}, "evbodi", NotFound},
		// Test with empty slice
		{[]string{}, "cheezaram", NotFound},
		// Test with nil slice
		{nil, "", NotFound},
	}

	for _, datum := range data {
		reval, err := ExponentialSearch(datum.slice, datum.value)
		if assert.NoError(err) {
			assert.Equal(datum.result, reval)
		}
	}
}
