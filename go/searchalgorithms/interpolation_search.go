package searchalgorithms

import (
	"fmt"
	"log"
)

// InterpolationSearch is a search algorithm that finds a value in a sorted slice of values.
//
// It takes a sorted slice of integers and the value to search for
// and returns the index of the first occurence of the value in the slice
// or “NotFound“ if the value is not found.
func InterpolationSearch(slice []int, value int) (int, error) {
	// Set up logging file
	close, err := setUpLogFile()
	if err != nil {
		return 0, fmt.Errorf("InterpolationSearch failed to open file: %w", err)
	}
	defer close()

	log.Printf("Performing Interpolation Search on slice %v\n", slice)

	l_bound := 0
	u_bound := len(slice) - 1

	// Use interpolation formula to find the index of the value
	for l_bound <= u_bound && slice[l_bound] <= value && slice[u_bound] >= value {
		pos := l_bound + ((u_bound-l_bound)*(value-slice[l_bound]))/(slice[u_bound]-slice[l_bound])
		log.Printf("Value at index [%d] = [%d]\n", pos, slice[pos])

		switch {
		case slice[pos] > value:
			u_bound = pos - 1
		case slice[pos] < value:
			l_bound = pos + 1
		default:
			// Make sure it is the first occurence
			for pos > 0 && slice[pos-1] == value {
				pos--
			}
			log.Printf("Value [%v] found at index [%v]\n\n", value, pos)
			return pos, nil
		}
	}

	// Value was not found
	log.Printf("Value [%v] not found in slice\n\n", value)
	return NotFound, nil
}
