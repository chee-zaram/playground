package searchalgorithms

import (
	"fmt"
	"log"
	"math"

	"golang.org/x/exp/constraints"
)

// JumpSearch uses jump search algorithm to find a value in a sorted slice.
//
// It takes in an Ordered sorted slice and a value and returns the index of the value
// and any errors.
func JumpSearch[T constraints.Ordered](slice []T, value T) (int, error) {
	// Set up logging in a temp file
	close, err := setUpLogFile()
	if err != nil {
		return 0, fmt.Errorf("JumpSearch failed to open file: %w", err)
	}
	defer close()

	log.Printf("Performing Jump Search on slice %v for value %v", slice, value)

	// Determine optimal jump size
	size := len(slice)
	step := int(math.Sqrt(float64(size)))
	log.Printf("Jump size is %d\n", step)

	// Find block where value may exist
	var u_bound, l_bound int
	for u_bound < size && slice[u_bound] < value {
		l_bound = u_bound
		u_bound += step
		log.Printf("Value at index [%d] = [%v]\n", l_bound, slice[l_bound])
	}

	// If block extends beyond the slice, make upper bound last index of slice
	if u_bound >= size {
		u_bound = size - 1
	}

	log.Printf("Value [%v] should be in block [%d,%d]. Performing linear search.\n", value, l_bound, u_bound)

	// Do linear search within block to find value
	for l_bound <= u_bound {
		if slice[l_bound] == value {
			log.Printf("Value [%v] found at index [%d]\n\n", value, l_bound)
			return l_bound, nil
		}
		l_bound++
	}

	log.Printf("Value [%v] not found in slice\n\n", value)
	return NotFound, nil
}
